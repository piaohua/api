package jtpay

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type JTpayTrans struct {
	Config *JTpayConfig
}

func NewJTpayTrans(cfg *JTpayConfig) (*JTpayTrans, error) {
	if cfg.Usercode == "" ||
		cfg.CompKey == "" ||
		cfg.NotifyUrl == "" ||
		cfg.ReturnUrl == "" ||
		cfg.PlaceOrderUrl == "" {
		return &JTpayTrans{Config: cfg}, errors.New("config field canot empty string")
	}
	return &JTpayTrans{Config: cfg}, nil
}

//交易结果通知验证
func (this *JTpayTrans) NotifyVerify(order *NotifyResult) bool {
	sign := p10sign(order.P1_usercode, order.P2_order,
		order.P3_money, order.P4_status, order.P5_jtpayorder,
		order.P6_paymethod, order.P8_charset,
		order.P9_signtype, this.Config.CompKey)
	//fmt.Println(sign, order.P10_sign)
	return order.P10_sign == sign
}

//打包post数据
func (this *JTpayTrans) Pack(order *JTpayOrder) string {
	//order.P7_sign = p7sign(this.Config.Usercode, order.P2_order,
	//	order.P3_money, order.P4_returnurl, order.P5_notifyurl,
	//	order.P6_ordertime, this.Config.CompKey)
	//order.P17_customip = customip(order.P17_customip)
	return this.compose(order)
}

//初始化订单
func (this *JTpayTrans) InitOrder(order *JTpayOrder) {
	order.P1_usercode = this.Config.Usercode
	order.P6_ordertime = OrderTimeStr()
	order.P2_order = GenOrderid(order.P1_usercode, order.P6_ordertime)
	order.P4_returnurl = this.Config.ReturnUrl
	order.P5_notifyurl = this.Config.NotifyUrl
	order.P14_customname = this.Config.Usercode
	//
	order.P7_sign = p7sign(this.Config.Usercode, order.P2_order,
		order.P3_money, order.P4_returnurl, order.P5_notifyurl,
		order.P6_ordertime, this.Config.CompKey)
	order.P17_customip = customip(order.P17_customip)
}

// 下单
func (this *JTpayTrans) Submit(order *JTpayOrder) (string, error) {
	//
	body := this.Pack(order)
	//fmt.Printf("body %s, order %#v\n", body, order)
	resp, err := doHttpPost(this.Config.PlaceOrderUrl, body)
	//fmt.Printf("resp %s, err %v", string(resp), err)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// 组装request报文
func (this *JTpayTrans) compose(order *JTpayOrder) string {
	val := make(url.Values)
	val.Set("p1_usercode", order.P1_usercode)
	val.Set("p2_order", order.P2_order)
	val.Set("p3_money", order.P3_money)
	val.Set("p4_returnurl", this.Config.ReturnUrl)
	val.Set("p5_notifyurl", this.Config.NotifyUrl)
	val.Set("p6_ordertime", order.P6_ordertime)
	val.Set("p7_sign", order.P7_sign)
	val.Set("p9_paymethod", order.P9_paymethod)
	val.Set("p14_customname", order.P14_customname)
	val.Set("p17_customip", order.P17_customip)
	val.Set("p25_terminal", order.P25_terminal)
	val.Set("p26_iswappay", order.P26_iswappay)
	return val.Encode()
}

// doRequest post the order in json format with a sign
func doHttpPost(targetUrl string, body string) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, strings.NewReader(body))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   0,
			KeepAlive: 0,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := &http.Client{Transport: transport}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}
