package wxsdk

import "encoding/json"

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func ParseResult(resp []byte) (Result, error) {
	result := Result{}
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

type CardsResult struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data CardsData `json:"data"`
}

type CardsData struct {
	Openid   string `json:"openid"`
	Card_num string `json:"card_num"`
}

func ParseCardsResult(resp []byte) (CardsResult, error) {
	result := CardsResult{}
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

type JsSdkResult struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data JsSdkData `json:"data"`
}

type JsSdkData struct {
	Code      string `json:"code"`
	Noncestr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
	Appid     string `json:"appid"`
}

func ParseJsSdkResult(resp []byte) (JsSdkResult, error) {
	jsSdkResult := JsSdkResult{}
	err := json.Unmarshal(resp, &jsSdkResult)
	if err != nil {
		return jsSdkResult, err
	}

	return jsSdkResult, nil
}
