# iapppay

Go语言iapp支付后台实现

Backend implementation of iapppay in golang 


# usage

```go
//初始化
cfg := &iapppay.IappConfig{
	AppId:         "应用编号,平台分配",
	PublicKeyPath: "应用公钥",
	PrivateKeyPath:"应用私钥",
	NotifyPath:    "支付结果回调路径",
	NotifyUrl:     "支付结果通知地址",
	PlaceOrderUrl: "http://ipay.iapppay.com:9999/payapi/order",
	QueryOrderUrl: "http://ipay.iapppay.com:9999/payapi/queryresult",
	Cipher:        rsa.Cipher,
}
appTrans, err := iapppay.NewAppTrans(cfg)
if err != nil {
	panic(err)
}

//获取transid(平台生成交易流水号),客户端得到orderid(应用内订单号)
transid, err := appTrans.Submit(price, orderid, uid, itemid)
if err != nil {
	panic(err)
}
fmt.Println(transid)

//查询订单接口
queryResult, err := appTrans.Query(orderid, uid)
if err != nil {
	panic(err)
}
fmt.Println(queryResult)

```

# document
