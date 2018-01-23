package wxapi

import (
	"testing"
)

var WxLogin *Wxapi //微信登录

func WxLoginInit() {
	cfg := &WxapiConfig{
		AppId:           "wx73057982c5a19e0c",
		AppSecret:       "82aeeb8c3ddb5a6c71103019bece6dbf",
		AccessUrl:       "https://api.weixin.qq.com/sns/oauth2/access_token",
		RefreshUrl:      "https://api.weixin.qq.com/sns/oauth2/refresh_token",
		UserinfoUrl:     "https://api.weixin.qq.com/sns/userinfo",
		VerifyAccessUrl: "https://api.weixin.qq.com/sns/auth",
	}
	wx, err := NewWxapi(cfg)
	if err != nil {
		panic(err)
	}
	WxLogin = wx
}

func TestRefresh(t *testing.T) {
	WxLoginInit()
	code := "021Odmeo0sPwHo1Nqxco0Glneo0OdmeH"
	accessResult, err := WxLogin.Auth(code)
	t.Log(err)
	t.Log(accessResult)
}
