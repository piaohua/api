package wxpay

type WxConfig struct {
	AppId         string
	AppKey        string
	MchId         string
	NotifyPattern string
	NotifyUrl     string
	PlaceOrderUrl string
	QueryOrderUrl string
	TradeType     string
}
