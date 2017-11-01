package wxapi

import "encoding/json"

type AccessResult struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
}

func ParseAccessResult(resp []byte) (AccessResult, error) {
	accessResult := AccessResult{}
	err := json.Unmarshal(resp, &accessResult)
	if err != nil {
		return accessResult, err
	}

	return accessResult, nil
}

type RefreshResult struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func ParseRefreshResult(resp []byte) (RefreshResult, error) {
	refreshResult := RefreshResult{}
	err := json.Unmarshal(resp, &refreshResult)
	if err != nil {
		return refreshResult, err
	}

	return refreshResult, nil
}

type UserInfoResult struct {
	ErrCode     int      `json:"errcode"`
	ErrMsg      string   `json:"errmsg"`
	OpenId      string   `json:"openid"`
	Nickname    string   `json:"nickname"`
	Sex         int      `json:"sex"`
	Province    string   `json:"province"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	HeadImagUrl string   `json:"headimgurl"`
	Privilege   []string `json:"privilege"`
	UnionId     string   `json:"unionid"`
}

func ParseUserInfoResult(resp []byte) (UserInfoResult, error) {
	userInfoResult := UserInfoResult{}
	err := json.Unmarshal(resp, &userInfoResult)
	if err != nil {
		return userInfoResult, err
	}

	return userInfoResult, nil
}

type VerifyAccessResult struct {
	ErrCode     int      `json:"errcode"`
	ErrMsg      string   `json:"errmsg"`
}

func ParseVerifyAccessResult(resp []byte) (VerifyAccessResult, error) {
	verifyAccessResult := VerifyAccessResult{}
	err := json.Unmarshal(resp, &verifyAccessResult)
	if err != nil {
		return verifyAccessResult, err
	}

	return verifyAccessResult, nil
}
