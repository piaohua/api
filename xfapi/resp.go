package xfapi

import "encoding/json"

// {
//     "ret": 0,
//     "token": "123456789",
//     "expire": "xxxx" // 过期时间，unix时间戳
// }
type GetTokenResult struct {
	Ret    int    `json:"ret"`
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}

func ParseGetTokenResult(resp []byte) (*GetTokenResult, error) {
	getTokenResult := new(GetTokenResult)
	err := json.Unmarshal(resp, getTokenResult)
	if err != nil {
		return nil, err
	}

	return getTokenResult, nil
}

// {
//     "ret": 0,
//     "token": "xxxx",
//     "expire": "xxxx", // 过期时间，unix时间戳
//     "sid": "abadafad@1234" //session id
// }
type GetUserTokenResult struct {
	Ret    int    `json:"ret"`
	Token  string `json:"token"`
	Expire int    `json:"expire"`
	Sid    string `json:"sid"`
}

func ParseGetUserTokenResult(resp []byte) (*GetUserTokenResult, error) {
	getUserTokenResult := new(GetUserTokenResult)
	err := json.Unmarshal(resp, getUserTokenResult)
	if err != nil {
		return nil, err
	}

	return getUserTokenResult, nil
}

// {
//     "ret":0,
//     "sid":"d3a12a9c-652b-4e94-95bb-50b27adadfaf",
//     "detail":"Succeed"
// }
//
// {
//     "ret"       :   0,
//     "gid"       :   "100001",
//     "detail"    :   "create group successfully"
// }
//
// {
//     "ret"       :   0,
//     "detail"    :   "Appoint managers successfully"
// }
type Result struct {
	Ret       int         `json:"ret"`
	Sid       string      `json:"sid"`
	Gid       string      `json:"gid"`
	Detail    string      `json:"detail"`
	Info      GroupInfo   `json:"info"`
	Grouplist []GroupList `json:"grouplist"`
}

func ParseResult(resp []byte) (*Result, error) {
	result := new(Result)
	err := json.Unmarshal(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GroupInfo struct {
	Appid        string   `json:"appid"`
	Gname        string   `json:"gname"`
	Gid          string   `json:"gid"`
	Members      []string `json:"members"`
	Type         int      `json:"type"`
	Owner        string   `json:"owner"`
	Managers     []string `json:"managers"`
	Maxusers     int      `json:"maxusers"`
	Announcement string   `json:"announcement"`
	Describe     string   `json:"describe"`
}

type GroupList struct {
	Owner    string `json:"owner"`
	Gname    string `json:"gname"`
	Maxusers int    `json:"maxusers"`
	Gid      string `json:"gid"`
	Size     int    `json:"size"`
	Type     int    `json:"type"`
}
