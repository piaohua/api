package wxsdk

type WxsdkConfig struct {
	AppKey        string //应用密钥
	CreateRoomUrl string //创建房间调用接口
	CancelRoomUrl string //返还房卡调用接口
	GetCardsUrl   string //查询房卡调用接口
	JsSdkUrl      string //获取微信JSSDK配置信息接口
}
