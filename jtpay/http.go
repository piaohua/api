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

//JTpayTrans 竣付通支付
type JTpayTrans struct {
	Config *JTpayConfig
}

//NewJTpayTrans 初始化竣付通支付
func NewJTpayTrans(cfg *JTpayConfig) (*JTpayTrans, error) {
	if cfg.Appid == "" ||
		cfg.CompKey == "" ||
		cfg.PlaceOrderUrl == "" {
		return &JTpayTrans{Config: cfg}, errors.New("config field canot empty string")
	}
	return &JTpayTrans{Config: cfg}, nil
}

//NotifyVerify 交易结果通知验证
func (this *JTpayTrans) NotifyVerify(order *NotifyResult) bool {
	if order.P1_yingyongnum != this.Config.Appid {
		return false
	}
	sign := p10sign(order, this.Config.CompKey)
	//fmt.Println(sign, order.P10_sign)
	return order.P10_sign == sign
}

//InitOrder 初始化订单
func (this *JTpayTrans) InitOrder(order *JTpayOrder) {
	order.P1_yingyongnum = this.Config.Appid
	order.P6_ordertime = OrderTimeStr()
	order.P2_ordernumber = GenOrderid(order.P1_yingyongnum, order.P6_ordertime)
	order.P8_sign = p8sign(order, this.Config.CompKey)
	order.P16_customip = customip(order.P16_customip)
}

//Submit 下单
func (this *JTpayTrans) Submit(order *JTpayOrder) (string, error) {
	body := this.compose(order) //打包post数据
	//fmt.Printf("body %s, order %#v\n", body, order)
	resp, err := doHttpPost(this.Config.PlaceOrderUrl, body)
	//fmt.Printf("resp %s, err %v", string(resp), err)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// 组装request报文
func (this *JTpayTrans) compose(o *JTpayOrder) string {
	val := make(url.Values)
	val.Set("p1_yingyongnum", o.P1_yingyongnum)
	val.Set("p2_ordernumber", o.P2_ordernumber)
	val.Set("p3_money", o.P3_money)
	val.Set("p6_ordertime", o.P6_ordertime)
	val.Set("p7_productcode", o.P7_productcode)
	val.Set("p8_sign", o.P8_sign)
	val.Set("p9_signtype", o.P9_signtype)
	val.Set("p10_bank_card_code", o.P10_bank_card_code)
	val.Set("p11_cardtype", o.P11_cardtype)
	val.Set("p12_channel", o.P12_channel)
	val.Set("p13_orderfailertime", o.P13_orderfailertime)
	val.Set("p14_customname", o.P14_customname)
	val.Set("p15_customcontact", o.P15_customcontact)
	val.Set("p16_customip", o.P16_customip)
	val.Set("p17_product", o.P17_product)
	val.Set("p18_productcat", o.P18_productcat)
	val.Set("p19_productnum", o.P19_productnum)
	val.Set("p20_pdesc", o.P20_pdesc)
	val.Set("p21_version", o.P21_version)
	val.Set("p22_sdkversion", o.P22_sdkversion)
	val.Set("p23_charset", o.P23_charset)
	val.Set("p24_remark", o.P24_remark)
	val.Set("p25_terminal", o.P25_terminal)
	val.Set("p26_ext1", o.P26_ext1)
	val.Set("p27_ext2", o.P27_ext2)
	val.Set("p28_ext3", o.P28_ext3)
	val.Set("p29_ext4", o.P29_ext4)
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
