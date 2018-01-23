package wxpay

import (
	//"encoding/xml"
	"encoding/json"
)

type PaymentRequest struct {
	//XMLName   xml.Name `xml:"xml"`
	AppId     string `xml:"appid",json:"appid"`
	PartnerId string `xml:"partnerid",json:"partnerid"`
	PrepayId  string `xml:"prepayid",json:"prepayid"`
	Package   string `xml:"package",json:"package"`
	NonceStr  string `xml:"noncestr",json:"noncestr"`
	Timestamp string `xml:"timestamp",json:"timestamp"`
	Sign      string `xml:"sign",json:"sign"`
}

func ToJson(req *PaymentRequest) ([]byte, error) {
	return json.Marshal(req)
}
