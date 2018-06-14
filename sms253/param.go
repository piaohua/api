package sms253

import (
	"encoding/json"
	"net/url"
)

func newParam(account, password, phone, msg string) ([]byte, error) {
	params := make(map[string]string)
	params["account"] = account
	params["password"] = password
	params["phone"] = phone
	params["msg"] = url.QueryEscape(msg)
	params["report"] = "false"
	bytesData, err := json.Marshal(params)
	return bytesData, err
}

//{
// "code" : "0", //状态码
// "msgId" : "17041010383624511", //消息Id
// "errorMsg" : "", //失败状态码说明（成功返回空）
// "time" : "20170410103836" //响应时间
//}

//SmsReponse reponse message
type SmsReponse struct {
	Code     string `json:"code"`
	MsgID    string `json:"msgId"`
	ErrorMsg string `json:"errorMsg"`
	Time     string `json:"time"`
}

// parse the reponse message
func parse(resp []byte) (*SmsReponse, error) {
	result := new(SmsReponse)
	err := json.Unmarshal(resp, result)
	if err != nil {
		return result, err
	}
	return result, nil
}
