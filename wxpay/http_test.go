package wxpay

import "testing"

var Apppay *AppTrans //微信支付

func WxPayInit() {
	host := "127.0.0.1"
	port := "8002"
	pattern := "/mahjong/wxpay/notice"
	notifyUrl := "https://" + host + ":" + port + pattern
	cfg := &WxConfig{
		AppId:         "wx730579825ca9e1c0",
		AppKey:        "F6D123S3XWC3E2FVDFW33E4DS6FMND67",
		MchId:         "1378900662",
		NotifyPattern: pattern,
		NotifyUrl:     notifyUrl,
		PlaceOrderUrl: "https://api.mch.weixin.qq.com/pay/unifiedorder",
		QueryOrderUrl: "https://api.mch.weixin.qq.com/pay/orderquery",
		TradeType:     "APP",
	}
	appTrans, err := NewAppTrans(cfg)
	if err != nil {
		panic(err)
	}
	Apppay = appTrans
}

func TestPay(t *testing.T) {
	WxPayInit()
	orderid := "kdkgkdkghfgdk"
	price := 600
	body := "1"
	ip := "119.29.24.17"
	transid, err := Apppay.Submit(orderid, float64(price), body, ip)
	t.Log(err)
	t.Log(transid)
	orderid = "590e031ffee2fa7b15c01686"
	queryResult, err := Apppay.Query(orderid)
	t.Log(err)
	t.Logf("%#v", queryResult)
}
