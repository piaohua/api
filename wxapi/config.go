package wxapi

type WxapiConfig struct {
	AppId       string //应用唯一标识
	AppSecret   string //应用密钥
	AccessUrl   string //授权请求地址
	RefreshUrl  string //刷新或续期地址
	UserinfoUrl string //获取用户个人信息地址
	VerifyAccessUrl string //检验授权凭证(access_token)是否有效
}
