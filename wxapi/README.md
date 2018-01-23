# wxlogin

Go语言微信App登录后台实现

Backend implementation of weixin login(app) in golang 


# usage

```go
//初始化
cfg := &wxapi.Wxapi{
	AppId:           "应用程序Id, 从https://open.weixin.qq.com上可以看得到",
	AppSecret:       "API密钥, 在 商户平台->账户设置->API安全 中设置",
	AccessUrl:       "https://api.weixin.qq.com/sns/oauth2/access_token",
	RefreshUrl:      "https://api.weixin.qq.com/sns/oauth2/refresh_token",
	UserinfoUrl:     "https://api.weixin.qq.com/sns/userinfo",
	VerifyAccessUrl: "https://api.weixin.qq.com/sns/auth",
}

wxLogin, err := wxapi.NewWxapi(cfg)
if err != nil {
	panic(err)
}

// 通过code获取access_token
AccessResult, err := wxLogin.Auth(code)
if err != nil {
	panic(err)
}
fmt.Println(AccessResult)

// 刷新access_token有效期
RefreshResult, err := wxLogin.Refresh(refresh_token)
if err != nil {
	panic(err)
}
fmt.Println(RefreshResult)

// 获取用户个人信息
UserInfoResult, err := wxLogin.UserInfo(access_token)
if err != nil {
	panic(err)
}
fmt.Println(UserInfoResult)

// 检验授权凭证（access_token）是否有效
err := wxLogin.VerifyAuth(access_token)
if err != nil {
	panic(err)
}
fmt.Println("ok")

```

# document

# TODO:优化
```go
type AccessReqest struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

type RefreshReqest struct {
	Appid        string `json:"appid"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type UserInfoReqest struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}

type AccessReqest struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}

const (
	redirectOauthURL      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoURL           = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	checkAccessTokenURL   = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
)

func getAccessTokenURL(AppId, AppSecret, code string) string {
	return fmt.Sprintf(accessTokenURL, AppId, AppSecret, code)
}

func getRefreshAccessTokenURL(AppId, refresh_token string) string {
	return fmt.Sprintf(refreshAccessTokenURL, AppId, refresh_token)
}

func getUserInfoURL(AppId, access_token string) string {
	return fmt.Sprintf(userInfoURL, access_token, AppId)
}

func getCheckAccessTokenURL(AppId, access_token string) string {
	return fmt.Sprintf(checkAccessTokenURL, access_token, AppId)
}
```
