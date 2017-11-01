package iapppay

// 下单请求参数
type IapppayOrder struct {
	Appid         string `json:"appid"  bson:"Appid"`                // 应用编号
	Waresid       uint32 `json:"waresid" bson:"Waresid"`             // 商品编号
	Waresname     string `json:"waresname" bson:"Waresname"`         // 商品名称
	Cporderid     string `json:"cporderid" bson:"Cporderid"`         // 商户订单号
	Price         uint32 `json:"price" bson:"Price"`                 // 支付金额
	Currency      string `json:"currency" bson:"Currency"`           // 货币类型
	Appuserid     string `json:"appuserid" bson:"Appuserid"`         // 用户在商户应用的唯一标识
	Cpprivateinfo string `json:"cpprivateinfo" bson:"Cpprivateinfo"` // 商户私有信息
	Notifyurl     string `json:"notifyurl" bson:"Notifyurl"`         // 支付结果通知地址
}

// 主动请求交易结果请求参数
type IapppayQuery struct {
	Appid     string `json:"appid" bson:"Appid"`         // 应用编号
	Cporderid string `json:"cporderid" bson:"Cporderid"` // 商户订单号
}
