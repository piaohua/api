package iapppay

import "github.com/89hmdys/toast/rsa"

type IappConfig struct {
	AppId          string //应用编号
	PublicKeyPath  string //公钥
	PrivateKeyPath string //私钥
	NotifyPattern  string //支付结果回调地址
	NotifyUrl      string //支付结果回调通知地址
	PlaceOrderUrl  string //下单请求地址
	QueryOrderUrl  string //查询请求地址
	Cipher         rsa.Cipher
}
