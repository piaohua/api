package wxsdk

import "testing"

var WxSDK *Wxsdk

func WxSdkInit() {
	cfg := &WxsdkConfig{
		AppKey:        "df2fef(DFS9832njf23R#@R@#",
		CreateRoomUrl: "http://localhost/index.php?c=Api&a=create_room",
		JsSdkUrl:      "http://localhost/index.php?c=Api&a=wx_config",
	}
	sdk, err := NewWxsdk(cfg)
	if err != nil {
		panic(err)
	}
	WxSDK = sdk
}

func TestSDK(t *testing.T) {
	WxSdkInit()
	url := "http://nn.18bn.cn/index.html"
	r, err := WxSDK.JsSdkInfo(url)
	t.Logf("%#v, err %v", r, err)
}

func TestVerify(t *testing.T) {
	WxSdkInit()
	openid := "oIIMzv-9XFFsJnzpV1H-2B52pXAw"
	nickname := "sss"
	headimgurl := "http://wx.qlogo.cn/mmopen/"
	var sex uint32 = 1
	var time uint32 = 1498543285
	sign := "f8021f80fc1383e65f5815f467ca1a32"

	//mysign := "7e814212928f13bd77f3a0de85d5fb6e"
	err := WxSDK.VerifyLogin(openid, nickname, headimgurl, sign, sex, time)
	t.Log(err)
}
