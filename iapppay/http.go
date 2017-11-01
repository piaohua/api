package iapppay

import (
	"bytes"
	cto "crypto"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/89hmdys/toast/crypto"
	"github.com/89hmdys/toast/rsa"
)

type IappTrans struct {
	Config *IappConfig
}

func NewIappTrans(cfg *IappConfig) (*IappTrans, error) {
	if cfg.AppId == "" ||
		cfg.PublicKeyPath == "" ||
		cfg.PrivateKeyPath == "" ||
		cfg.NotifyPattern == "" ||
		cfg.NotifyUrl == "" ||
		cfg.QueryOrderUrl == "" ||
		cfg.PlaceOrderUrl == "" {
		return &IappTrans{Config: cfg}, errors.New("config field canot empty string")
	}
	key, err := rsa.LoadKeyFromPEMFile(
		cfg.PublicKeyPath,
		cfg.PrivateKeyPath,
		rsa.ParsePKCS1Key)
	if err != nil {
		return &IappTrans{Config: cfg}, err
	}
	cipher, err := crypto.NewRSA(key)
	if err != nil {
		return &IappTrans{Config: cfg}, err
	}
	cfg.Cipher = cipher
	return &IappTrans{Config: cfg}, nil
}

// RSA签名
func (this *IappTrans) IpaySign(tran []byte) ([]byte, error) {
	return this.Config.Cipher.Sign(tran, cto.MD5)
}

// RSA验签
func (this *IappTrans) IpayVerify(tran []byte, sign []byte) error {
	return this.Config.Cipher.Verify(tran, sign, cto.MD5)
}

// 下单
func (this *IappTrans) Submit(price uint32, orderId, uid, itemid, name string) (string, error) {
	iappOdr := this.newIappayOrder(price, orderId, uid, itemid, name)
	body, err := this.composeOrderReq(iappOdr)
	if err != nil {
		return "", err
	}
	resp, err := doHttpPost(this.Config.PlaceOrderUrl, []byte(body))
	if err != nil {
		return "", err
	}
	transid, err := this.ParsePlaceOrderResp(resp)
	if err != nil {
		return "", err
	}
	return transid, nil
}

// 下单请求参数
func (this *IappTrans) newIappayOrder(price uint32, id, uid, itemid, name string) *IapppayOrder {
	return &IapppayOrder{
		Appid:         this.Config.AppId,     //应用编号
		Waresid:       6,                     //商品编号
		Waresname:     name,                  //商品名称
		Cporderid:     id,                    //商户订单号
		Price:         price,                 //支付金额Float(6,2)
		Currency:      "RMB",                 //货币类型
		Appuserid:     uid,                   //游戏内用户唯一id
		Cpprivateinfo: itemid,                // 商品id或其他
		Notifyurl:     this.Config.NotifyUrl, //支付结果通知地址
	}
}

// 组装request报文
func (this *IappTrans) composeOrderReq(i *IapppayOrder) (string, error) {
	tran, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	sign, err := this.IpaySign(tran)
	if err != nil {
		return "", err
	}
	t := url.QueryEscape(string(tran))
	s := url.QueryEscape(base64.StdEncoding.EncodeToString(sign))
	b := fmt.Sprintf("transdata=%s&sign=%s&signtype=RSA", t, s)
	return b, nil
}

// 主动请求 组装request报文
func (this *IappTrans) composeQueryReq(i *IapppayQuery) (string, error) {
	tran, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	sign, err := this.IpaySign(tran)
	if err != nil {
		return "", err
	}
	t := url.QueryEscape(string(tran))
	s := url.QueryEscape(base64.StdEncoding.EncodeToString(sign))
	b := fmt.Sprintf("transdata=%s&sign=%s&signtype=RSA", t, s)
	return b, nil
}

// 下单请求参数
func (this *IappTrans) newIappayQuery(id string) *IapppayQuery {
	return &IapppayQuery{
		Appid:     this.Config.AppId, // 应用编号
		Cporderid: id,                // 商户订单号
	}
}

// 交易结果主动查询
// cporderid = data.ChargeOrder.Orderid,
// cporderid: 89, Transid: 32471610191058365861
func (this *IappTrans) Query(id, userid string) (*QueryResult, error) {
	iappQuery := this.newIappayQuery(id)
	body, err := this.composeQueryReq(iappQuery)
	if err != nil {
		return nil, err
	}
	resp, err := doHttpPost(this.Config.QueryOrderUrl, []byte(body))
	if err != nil {
		return nil, err
	}
	return this.ParseQueryResp(resp)
}

// 接收交易结果通知
func (this *IappTrans) RecvNotify(recv func(http.ResponseWriter, *http.Request)) {
	http.Handle(this.Config.NotifyPattern, http.HandlerFunc(recv))
}

// doRequest post the order in json format with a sign
func doHttpPost(targetUrl string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

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
