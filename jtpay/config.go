package jtpay

type JTpayConfig struct {
	Usercode      string //商户号
	CompKey       string //秘钥
	NotifyUrl     string //支付结果回调通知地址
	ReturnUrl     string //支付结果回调地址(用于告知付款人支付结果)
	PlaceOrderUrl string //下单请求地址
}
