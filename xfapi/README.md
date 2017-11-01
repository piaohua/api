# xfapi

Go语言xunfei api实现

Backend implementation of xunfei in golang

# usage

```go
//初始化
host := "http://im.voicecloud.cn:1208/v1"
cfg := &xfapi.XfapiConfig{
	AppId:              "99999999",
	AppSecret:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	AppToken:           "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
	GetTokenUrl:        host + "/rest/getToken.do",
	GetUserTokenUrl:    host + "/rest/getUserToken.do",
	UpdateUserTokenUrl: host + "/rest/updateUserToken.do",
	GroupCreateUrl:     host + "/group/create.do",
	GroupAddUrl:        host + "/group/add.do",
	GroupExitUrl:       host + "/group/exit.do",
}
xf, err := xfapi.NewXfapi(cfg)
if err != nil {
	panic(err)
}

```

# document

