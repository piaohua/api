package apple

import (
	"encoding/json"
)

// 苹果验证返回码对应信息
/*
Status 	Description
0 	The receipt provided is valid.
21000 	The App Store could not read the JSON object you provided.
21002 	The data in the receipt-data property was malformed.
21003 	The receipt could not be authenticated.
21004 	The shared secret you provided does not match the shared secret on file for your account.Only returned for iOS 6 style transaction receipts for auto-renewable subscriptions.
21005 	The receipt server is not currently available.
21006 	This receipt is valid but the subscription has expired. When this status code is returned to your server, the receipt data is also decoded and returned as part of the response.Only returned for iOS 6 style transaction receipts for auto-renewable subscriptions.
21007 	This receipt is a sandbox receipt, but it was sent to the production server.
21008 	This receipt is a production receipt, but it was sent to the sandbox server.

键名 描述
quantity 购买商品的数量。对应SKPayment对象中的quantity属性
product_id 商品的标识，对应SKPayment对象的productIdentifier属性。
transaction_id 交易的标识，对应SKPaymentTransaction的transactionIdentifier属性
purchase_date 交易的日期，对应SKPaymentTransaction的transactionDate属性
original_-transaction_id 对于恢复的transaction对象，该键对应了原始的transaction标识
original_purchase_-date 对于恢复的transaction对象，该键对应了原始的交易日期
app_item_id App Store用来标识程序的字符串。一个服务器可能需要支持多个server的支付功能，可以用这个标识来区分程序。链接sandbox用来测试的程序的不到这个值，因此该键不存在。
version_external_-identifier 用来标识程序修订数。该键在sandbox环境下不存在
bid iPhone程序的bundle标识
bvrs iPhone程序的版本号

"environment":"Sandbox",
"environment":"Production",

*/

// 苹果验证成功返回数据结构
/*
{
	"status":0,
	"environment":"Sandbox",
	"receipt":{
		"receipt_type":"ProductionSandbox",
		"adam_id":0,
		"app_item_id":0,
		"bundle_id":"gzmj01.ysgame.com",
		"application_version":"1.0",
		"download_id":0,
		"version_external_identifier":0,
		"receipt_creation_date":"2016-09-09 02:47:38 Etc/GMT",
		"receipt_creation_date_ms":"1473389258000",
		"receipt_creation_date_pst":"2016-09-08 19:47:38 America/Los_Angeles",
		"request_date":"2016-09-09 02:47:40 Etc/GMT",
		"request_date_ms":"1473389260393",
		"request_date_pst":"2016-09-08 19:47:40 America/Los_Angeles",
		"original_purchase_date":"2013-08-01 07:00:00 Etc/GMT",
		"original_purchase_date_ms":"1375340400000",
		"original_purchase_date_pst":"2013-08-01 00:00:00 America/Los_Angeles",
		"original_application_version":"1.0",
		"in_app":[
		{
			"quantity":"1",
			"product_id":"1",
			"transaction_id":"1000000234931388",
			"original_transaction_id":"1000000234931388",
			"purchase_date":"2016-09-09 02:47:38 Etc/GMT",
			"purchase_date_ms":"1473389258000",
			"purchase_date_pst":"2016-09-08 19:47:38 America/Los_Angeles",
			"original_purchase_date":"2016-09-09 02:47:38 Etc/GMT",
			"original_purchase_date_ms":"1473389258000",
			"original_purchase_date_pst":"2016-09-08 19:47:38 America/Los_Angeles",
			"is_trial_period":"false"
		}
		]
	}
}

&apple.TradeResult{Status:0, Environment:"Sandbox", Receipt:apple.Receipt{Receipt_type:"ProductionSandbox", Adam_id:0, App_item_id:0, Bundle_id:"lxqc.qpniuniu.com", Application_version:"1.0", Download_id:0, Version_external_identifier:0, Receipt_creation_date:"2017-06-03 11:40:58 Etc/GMT", Receipt_creation_date_ms:"1496490058000", Receipt_creation_date_pst:"2017-06-03 04:40:58 America/Los_Angeles", Request_date:"2017-06-03 11:42:51 Etc/GMT", Request_date_ms:"1496490171919", Request_date_pst:"2017-06-03 04:42:51 America/Los_Angeles", Original_purchase_date:"2013-08-01 07:00:00 Etc/GMT", Original_purchase_date_ms:"1375340400000", Original_purchase_date_pst:"2013-08-01 00:00:00 America/Los_Angeles", Original_application_version:"1.0", InApp:[]apple.InApp{apple.InApp{Quantity:"1", Product_id:"2", Transaction_id:"1000000304288600", Original_transaction_id:"", Purchase_date:"", Purchase_date_ms:"1496490058000", Purchase_date_pst:"2017-06-03 04:40:58 America/Los_Angeles", Original_purchase_date:"2017-06-03 11:40:58 Etc/GMT", Original_purchase_date_ms:"1496490058000", original_purchase_date_pst:"", Is_trial_periodstring:"false"}}}}

*/

// APPLE PAY URL
const (
	APPLE_VERIY_PRODUCTION_URL = "https://buy.itunes.apple.com/verifyReceipt"     // 正式验证地址
	APPLE_VERIY_SANDBOX_URL    = "https://sandbox.itunes.apple.com/verifyReceipt" //测试验证地址
)

type InApp struct {
	Quantity                   string `json:"quantity"`
	Product_id                 string `json:"product_id"`
	Transaction_id             string `json:"transaction_id"`
	Original_transaction_id    string `json:"original_transaction_id"`
	Purchase_date              string `json:"original_transaction_id"`
	Purchase_date_ms           string `json:"purchase_date_ms"`
	Purchase_date_pst          string `json:"purchase_date_pst"`
	Original_purchase_date     string `json:"original_purchase_date"`
	Original_purchase_date_ms  string `json:"original_purchase_date_ms"`
	original_purchase_date_pst string `json:"original_purchase_date_pst"`
	Is_trial_periodstring      string `json:"is_trial_period"`
}

type Receipt struct {
	Receipt_type                 string  `json:"receipt_type"`
	Adam_id                      int     `json:"adam_id"`
	App_item_id                  int     `json:"app_item_id"`
	Bundle_id                    string  `json:"bundle_id"`
	Application_version          string  `json:"application_version"`
	Download_id                  int     `json:"download_id"`
	Version_external_identifier  int     `json:"version_external_identifier"`
	Receipt_creation_date        string  `json:"Receipt_creation_date"`
	Receipt_creation_date_ms     string  `json:"receipt_creation_date_ms"`
	Receipt_creation_date_pst    string  `json:"receipt_creation_date_pst"`
	Request_date                 string  `json:"request_date"`
	Request_date_ms              string  `json:"request_date_ms"`
	Request_date_pst             string  `json:"request_date_pst"`
	Original_purchase_date       string  `json:"original_purchase_date"`
	Original_purchase_date_ms    string  `json:"original_purchase_date_ms"`
	Original_purchase_date_pst   string  `json:"original_purchase_date_pst"`
	Original_application_version string  `json:"original_application_version"`
	InApp                        []InApp `json:"in_app"`
}

// 苹果验证返回数据结构
type TradeResult struct {
	Status      int     `json:"status"`
	Environment string  `json:"environment"`
	Receipt     Receipt `json:"receipt"`
}

// Parse the reponse message
func ParseTradeResult(resp []byte) (*TradeResult, error) {
	result := new(TradeResult)
	err := json.Unmarshal(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// POST方式提交到苹果支付验证服务器的json数据结构
type Requst struct {
	ReceiptData string `json:"receipt-data"`
}

func ToJson(receipt string) ([]byte, error) {
	request := new(Requst)
	request.ReceiptData = receipt
	return json.Marshal(request)
}
