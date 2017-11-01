package iapppay

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// 请求失败结果
type FailMsg struct {
	Code   json.Number `json:"code,Number" bson:"Code"`
	Errmsg string      `json:"errmsg" bson:"Errmsg"`
}

// 下单请求成功结果
type PlaceOrderResult struct {
	Transid string `json:"transid" bson:"Transid"`
}

// 主动请求交易结果
type QueryResult struct {
	Cporderid string  `json:"cporderid" bson:"Cporderid"` // 商户订单号
	Transid   string  `json:"transid" bson:"Transid"`     // 交易流水号
	Appuserid string  `json:"appuserid" bson:"Appuserid"` // 用户在商户应用中唯一id
	Appid     string  `json:"appid" bson:"Appid"`         // 游戏id
	Waresid   uint32  `json:"waresid" bson:"Waresid"`     // 商品编码
	Feetype   int     `json:"feetype" bson:"Feetype"`     // 计费方式
	Money     float32 `json:"money" bson:"Money"`         // 交易金额
	Currency  string  `json:"currency" bson:"Currency"`   // 货币类型	RMB
	Result    int     `json:"result" bson:"Result"`       // 交易结果	0–交易成功 1–交易失败
	Transtime string  `json:"transtime" bson:"Transtime"` // 交易完成时间 yyyy-mm-dd hh24:mi:ss
	Cpprivate string  `json:"cpprivate" bson:"Cpprivate"` // 商户私有信息
	Paytype   uint32  `json:"paytype" bson:"Paytype"`     // 支付方式
}

// 交易结果通知
type TradeResult struct {
	Transtype int     `json:"transtype"` // 交易类型0–支付交易；
	Cporderid string  `json:"cporderid"` // 商户订单号
	Transid   string  `json:"transid"`   // 交易流水号
	Appuserid string  `json:"appuserid"` // 用户在商户应用中唯一id
	Appid     string  `json:"appid"`     // 游戏id
	Waresid   uint32  `json:"waresid"`   // 商品编码
	Feetype   int     `json:"feetype"`   // 计费方式
	Money     float32 `json:"money"`     // 交易金额
	Currency  string  `json:"currency"`  // 货币类型	RMB
	Result    int     `json:"result"`    // 交易结果	0–交易成功 1–交易失败
	Transtime string  `json:"transtime"` // 交易完成时间 yyyy-mm-dd hh24:mi:ss
	Cpprivate string  `json:"cpprivate"` // 商户私有信息
	Paytype   uint32  `json:"paytype"`   // 支付方式
}

// 解析response报文, 返回transid
// transdata={"transid":"11111"}&sign=xxxxxx&signtype=RSA
// transdata={"code":"1001","errmsg":"签名验证失败"}
func (this *IappTrans) ParsePlaceOrderResp(result []byte) (string, error) {
	m := ParseResponse(result)
	if this.ParseVerify(m) {
		var r PlaceOrderResult
		err := json.Unmarshal([]byte(m["transdata"]), &r)
		if err != nil {
			return "", err
		}
		return r.Transid, nil
	}
	var f FailMsg
	err := json.Unmarshal([]byte(m["transdata"]), &f)
	if err != nil {
		return "", err
	}
	return "", fmt.Errorf("return Errmsg: %s, Code: %v", f.Errmsg, f.Code)
}

func (this *IappTrans) ParseVerify(m map[string]string) bool {
	if sign, ok := m["sign"]; ok {
		bufs, err := base64.StdEncoding.DecodeString(sign)
		if err != nil {
			return false
		}
		err = this.IpayVerify([]byte(m["transdata"]), bufs)
		if err == nil && m["signtype"] == "RSA" {
			return true
		}
	}
	return false
}

// 解析response报文
func (this *IappTrans) ParseQueryResp(result []byte) (*QueryResult, error) {
	m := ParseResponse(result)
	if this.ParseVerify(m) {
		var q QueryResult
		err := json.Unmarshal([]byte(m["transdata"]), &q)
		if err != nil {
			return nil, err
		}
		return &q, nil
	}
	var f FailMsg
	err := json.Unmarshal([]byte(m["transdata"]), &f)
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("return Errmsg: %s, Code: %v", f.Errmsg, f.Code)
}

// 接收交易结果通知
func (this *IappTrans) ParseTradeResult(result []byte) (*TradeResult, error) {
	m := ParseResponse(result)
	if this.ParseVerify(m) {
		//解析返回的JSON数据
		var t TradeResult
		err := json.Unmarshal([]byte(m["transdata"]), &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	return nil, fmt.Errorf("return sign err: %v", m)
}

func ParseResponse(result []byte) map[string]string {
	m := make(map[string]string)
	r := strings.Split(string(result), "&")
	for _, v := range r {
		h := strings.Split(v, "=")
		// m[h[0]] = h[1]
		arg, err := url.QueryUnescape(h[1])
		if err != nil {
			m[h[0]] = ""
			continue
		}
		m[h[0]] = arg
	}
	return m
}
