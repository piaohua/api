package feiyu

type FeiyuConfig struct {
	//'配置项'   '配置值'
	//账号信息
	USER_INFO string
	//绑定号码
	USER_PHONE string
	//账号密码
	USER_PASSWORD string
	//测试APPID
	APP_ID string
	//测试APPTOKEN
	APP_TOKEN string
	//当前接口的版本号
	VERSION string
	//添加终端SDK账户
	ADD_ACCOUNT_URL string //"http://api.feiyucloud.com/api/addAccount",
	//查看终端SDK账户
	GET_ACCOUNT_URL string //"http://api.feiyucloud.com/api/getAccount",
	//屏蔽终端SDK账户
	DISABLE_ACCOUNT_URL string //"http://api.feiyucloud.com/api/disableAccount",
	//激活终端SDK账户
	ENABLE_ACCOUNT_URL string //"http://api.feiyucloud.com/api/enableAccount",
	//修改终端SDK账户绑定手机号
	MODIFY_ACCOUNT_DISPLAY_NUMBER_URL string //"http://api.feiyucloud.com/api/modifyAccountDisplayNumber",
	//查询SDK账户的在线状态
	GET_ACCOUNT_ONLINE_STATUS_URL string //"http://api.feiyucloud.com/api/onlineStatus",
	//回拨（集成方服务器端主动发起回拨）
	CALL_BACK_URL string //"http://api.feiyucloud.com/api/callback",
	//获取录音文件下载地址
	GET_RECORD_DOWN_URL string //"http://api.feiyucloud.com/api/getRecordDownUrl",
	//语音通知
	SOUND_SMS_ACTION string //"http://soundsms.feiyucloud.com/soundSmsAction!autocall.action",
	//查询通话记录
	FETCH_CALL_HISTORY string //"http://api.feiyucloud.com/api/fetchCallHistory"
	//播放语音
	CONFERENCE_PLAY string //http://confapi.feiyucloud.com/conference/play
	//取消静音
	CONFERENCE_UNMUTE string //http://confapi.feiyucloud.com/conference/unmute
	//静音
	CONFERENCE_MUTE string //http://confapi.feiyucloud.com/conference/mute
	//挂断
	CONFERENCE_HANGUP string //http://confapi.feiyucloud.com/conference/hangup
	//话单推送
	TEL_LIST_PUSH string //http://localhost:7015/mahjong/feiyu/telListPush
	//呼叫授权
	CALL_AUTH string //http://localhost:7015/mahjong/feiyu/callAuth
}

func NewFeiyuConfig() *FeiyuConfig {
	fy := &FeiyuConfig{
		//账号信息
		USER_INFO: "XXX",
		//绑定号码
		USER_PHONE: "XXX",
		//账号密码
		USER_PASSWORD: "XXX",
		//测试APPID
		APP_ID: "C814FB316B922EE6D55572F9FCAF4735",
		//测试APPTOKEN
		APP_TOKEN: "86ACAF54093030CEA6E866C105A93126",
		//当前接口的版本号
		VERSION: "2.1.0",
		//添加终端SDK账户
		ADD_ACCOUNT_URL: "http://api.feiyucloud.com/api/addAccount",
		//查看终端SDK账户
		GET_ACCOUNT_URL: "http://api.feiyucloud.com/api/getAccount",
		//屏蔽终端SDK账户
		DISABLE_ACCOUNT_URL: "http://api.feiyucloud.com/api/disableAccount",
		//激活终端SDK账户
		ENABLE_ACCOUNT_URL: "http://api.feiyucloud.com/api/enableAccount",
		//修改终端SDK账户绑定手机号
		MODIFY_ACCOUNT_DISPLAY_NUMBER_URL: "http://api.feiyucloud.com/api/modifyAccountDisplayNumber",
		//查询SDK账户的在线状态
		GET_ACCOUNT_ONLINE_STATUS_URL: "http://api.feiyucloud.com/api/onlineStatus",
		//回拨（集成方服务器端主动发起回拨）
		CALL_BACK_URL: "http://api.feiyucloud.com/api/callback",
		//获取录音文件下载地址
		GET_RECORD_DOWN_URL: "http://api.feiyucloud.com/api/getRecordDownUrl",
		//语音通知
		SOUND_SMS_ACTION: "http://soundsms.feiyucloud.com/soundSmsAction!autocall.action",
		//查询通话记录
		FETCH_CALL_HISTORY: "http://api.feiyucloud.com/api/fetchCallHistory",
		//播放语音
		CONFERENCE_PLAY: "http://confapi.feiyucloud.com/conference/play",
		//取消静音
		CONFERENCE_UNMUTE: "http://confapi.feiyucloud.com/conference/unmute",
		//静音
		CONFERENCE_MUTE: "http://confapi.feiyucloud.com/conference/mute",
		//挂断
		CONFERENCE_HANGUP: "http://confapi.feiyucloud.com/conference/hangup",
		//话单推送
		TEL_LIST_PUSH: "/mahjong/feiyu/telListPush",
		//呼叫授权
		CALL_AUTH: "/mahjong/feiyu/callAuth",
	}
	return fy
}
