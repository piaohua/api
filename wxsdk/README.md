# wxsdk

# usage

```go
//初始化

var WxSDK *Wxsdk

func WxSdkInit() {
	cfg := &WxsdkConfig{
		AppKey:        "xx",
		CreateRoomUrl: "http://localhost/index.php?c=Api&a=create_room",
		JsSdkUrl:      "http://localhost/index.php?c=Api&a=wx_config",
	}
	sdk, err := NewWxsdk(cfg)
	if err != nil {
		panic(err)
	}
	WxSDK = sdk
}

r, err := WxSDK.JsSdkInfo(url)
fmt.Printf("r %v, err %v", r, err)

```

# document

一、创建房间调用接口
1.地址 http://localhost/index.php?c=Api&a=create_room
2.请求方式建议使用get
3.参数：
	game 		牛牛游戏直接传1
	room_id  	房间编号
	openid		用户的openid，平台跳转过去时会传过去
	card_num	本次创建房间需要消耗的房卡数量
	time		服务器当前时间戳 单位为秒
	sign		参考签名规则
4.返回值json格式 -1是未成功，具体错误参考返回的msg，1是成功
{"code":"1","msg":"创建成功"}


二、获取微信JSSDK配置信息接口
1.地址 http://localhost/index.php?c=Api&a=wx_config
2.请求方式建议使用get
3.参数：
	url			当前需要签名页面的地址，请先进行urlencode
	time		服务器当前时间戳 单位为秒
	sign		参考签名规则
4.返回值json格式 -1是未成功，具体错误参考返回的msg，1是成功
{"code":"1","msg":"ok","data":{"noncestr":"noncestr_289544","timestamp":1498282189,"url":"http:\/\/www.baidu.com","signature":"27739858ad1716dcf3c39f4ce359f4cf98cbe3e4","appid":"wxae836f7e27b91660"}}

三、平台跳转进入游戏接口（此接口为平台跳转过去游戏实例接收）
1.地址 http://localhost:81/index.html?time=1498281098&openid=oIIMzv0GKfGAqu3xA7E-HvIjYAX8&nickname=%E8%84%9A%E6%9C%AC&sex=1&headimgurl=http%253A%252F%252Fwx.qlogo.cn%252Fmmopen%252F52dvIll9rc9oiay2Jz3zDM6vAicgPz79DoB9Cqz5ZPBiadt09icSqAzqia4dK3PZia2wqVw0MAY7MqAkY3TXdNmBQJzPjxjJP7kzIg%252F0&sign=6ae37fa42ce8bcc582f72ee76de40176

2.参数：
	openid		当前登录用户的openid
	nickname	微信用户昵称
	sex			性别
	headimgurl	微信用户头像地址
	time		服务器当前时间戳 单位为秒
	sign		参考签名规则，游戏端在接收参数时也请进行签名校验，且根据time参数限定接口地址的有效时间，建议3分钟内有效

四、签名规则
1.所有传入参数按照字典序进行排序，拼接字符串最后连接签名key进行md5加密，加密结果转成小写
2.签名使用key   df2fef(DFS9832njf23R#@R@#
3.实例:
	card_num=2&game=1&openid=oIIMzv0GKfGAqu3xA7E-HvIjYAX8&room_id=12580&time=1498281583&key=df2fef(DFS9832njf23R#@R@#
	将上面的字符串进行md5，转换成小写即可得到$sign参数
