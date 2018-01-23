package wxapi

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Wxapi is abstact of weixin api handler.
type Wxapi struct {
	Config *WxapiConfig
}

// Initialized the AppTrans with specific config
func NewWxapi(cfg *WxapiConfig) (*Wxapi, error) {
	if cfg.AppId == "" ||
		cfg.AppSecret == "" ||
		cfg.AccessUrl == "" ||
		cfg.RefreshUrl == "" ||
		cfg.UserinfoUrl == "" {
		return &Wxapi{Config: cfg}, errors.New("config field canot empty string")
	}

	return &Wxapi{Config: cfg}, nil
}

// 通过code获取access_token
func (this *Wxapi) Auth(code string) (AccessResult, error) {
	accessResult := AccessResult{}

	queryUrl := this.newAccReqParam(code)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return accessResult, err
	}

	accessResult, err = ParseAccessResult(resp)
	if err != nil {
		return accessResult, err
	}

	//verity ErrCode of response
	if accessResult.ErrCode != 0 {
		return accessResult, fmt.Errorf("return ErrCode:%d, ErrMsg:%s", accessResult.ErrCode, accessResult.ErrMsg)
	}

	return accessResult, nil
}

func (this *Wxapi) newAccReqParam(code string) string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["secret"] = this.Config.AppSecret
	param["code"] = code
	param["grant_type"] = "authorization_code"

	return this.Config.AccessUrl + ToQueryString(param)
}

// 刷新access_token有效期
func (this *Wxapi) Refresh(refreshToken string) (RefreshResult, error) {
	refreshResult := RefreshResult{}

	queryUrl := this.newRefReqParam(refreshToken)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return refreshResult, err
	}

	refreshResult, err = ParseRefreshResult(resp)
	if err != nil {
		return refreshResult, err
	}

	//verity ErrCode of response
	if refreshResult.ErrCode != 0 {
		return refreshResult, fmt.Errorf("return ErrCode:%d, ErrMsg:%s", refreshResult.ErrCode, refreshResult.ErrMsg)
	}

	return refreshResult, nil
}

func (this *Wxapi) newRefReqParam(refreshToken string) string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["refresh_token"] = refreshToken
	param["grant_type"] = "refresh_token"

	return this.Config.RefreshUrl + ToQueryString(param)
}

// 获取用户个人信息
func (this *Wxapi) UserInfo(openid, accessToken string) (UserInfoResult, error) {
	userInfoResult := UserInfoResult{}

	queryUrl := this.newUserReqParam(openid, accessToken)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return userInfoResult, err
	}

	userInfoResult, err = ParseUserInfoResult(resp)
	if err != nil {
		return userInfoResult, err
	}

	//verity ErrCode of response
	if userInfoResult.ErrCode != 0 {
		return userInfoResult, fmt.Errorf("return ErrCode:%d, ErrMsg:%s", userInfoResult.ErrCode, userInfoResult.ErrMsg)
	}

	return userInfoResult, nil
}

func (this *Wxapi) newUserReqParam(openid, accessToken string) string {
	param := make(map[string]string)
	param["openid"] = openid
	param["access_token"] = accessToken
	param["lang"] = "zh-CN"

	return this.Config.UserinfoUrl + ToQueryString(param)
}

// 检验授权凭证（access_token）是否有效
func (this *Wxapi) VerifyAuth(openid, accessToken string) error {

	queryUrl := this.newVerifyReqParam(openid, accessToken)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return err
	}

	verifyAccessResult, err := ParseVerifyAccessResult(resp)
	if err != nil {
		return err
	}

	//verity ErrCode of response
	if verifyAccessResult.ErrCode != 0 {
		return fmt.Errorf("return ErrCode:%d, ErrMsg:%s", verifyAccessResult.ErrCode, verifyAccessResult.ErrMsg)
	}

	return nil
}

func (this *Wxapi) newVerifyReqParam(openid, accessToken string) string {
	param := make(map[string]string)
	param["openid"] = openid
	param["access_token"] = accessToken

	return this.Config.VerifyAccessUrl + ToQueryString(param)
}

// doRequest get the order in json format with a sign
func doHttpPost(targetUrl string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("GET", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}
