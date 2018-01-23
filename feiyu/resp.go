package feiyu

import (
	"encoding/json"
	"strconv"
	"strings"
)

//往飞语服务器添加飞语云账户的响应结果
/**
 * result：返回的具体数据，当resultCode为0表示成功，result存储的是由以下几个参数构成的json值。
 * 返回参数说明：
 * fyAccountId：String必有，飞语云账户ID，终端授权需要
 * fyAccountPwd：String必有，飞语云账户密码，终端授权需要
 * createDate：String必有，创建时间。格式：yyyy-MM-dd HH:mm:ss
 * status：int必有，状态。1）有效；2）被屏蔽
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回数据样例：{"result":{"addDate":"2015-03-17 14:06:19","fyAccountId":"123123","fyAccountPwd":"1223456","status":1}, "resultCode":"0","resultMsg":"创建账户成功"}
 */
type AddAccountResult struct {
	Result ResultInfo `json:"result"` //
	ResultCode	string `json:"resultCode"`  //	是	返回的错误代码,0）代表成功；其他具体的错误代码见错误描述
	ResultMsg	string `json:"resultMsg"`  //	否	错误信息描述
}

type ResultInfo struct {
	FyAccountId	string `json:"fyAccountId"`  //	是	飞语云账户ID
	FyAccountPwd	string `json:"fyAccountPwd"`  //	是	飞语云账户密码
	CreateDate	string `json:"addDate"`  //	是	创建时间,格式：yyyy-MM-dd HH:mm:ss
	Status	int	`json:"status"` //是	状态,1）启用；2）禁用
}

//解析响应结果
func ParseAddAccountResult(resp []byte) (AddAccountResult, error) {
	addAccountResult := AddAccountResult{}
	err := json.Unmarshal(resp, &addAccountResult)
	if err != nil {
		return addAccountResult, err
	}

	return addAccountResult, nil
}

//往飞语服务器查询飞语云账户的响应结果
/**
 * result：
 * 返回的具体数据，当resultCode为0表示成功，result存储的是由以下几个参数构成的json值。
 返回参数说明：
 fyAccountId：String必有，飞语云账户ID，终端授权需要
 fyAccountPwd：String必有，飞语云账户密码，终端授权需要
 createDate：String必有，创建时间。格式：yyyy-MM-dd HH:mm:ss
 status：int必有，状态。1）有效；2）被屏蔽
 resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 resultMsg：错误信息描述
 返回数据样例：{
 	"result": {
 		"addDate": "2015-03-17 14:06:19",
 		"fyAccountId": "123123",
 		"fyAccountPwd": "1223456",
 		"status": 1
 	},
 	"resultCode": "0",
 	"resultMsg": ""
 }
 */
type GetAccountResult struct {
	Result ResultInfo `json:"result"` //
	ResultCode	string `json:"resultCode"`  //	是	返回的错误代码,0）代表成功；其他具体的错误代码见错误描述
	ResultMsg	string `json:"resultMsg"`  //	否	错误信息描述
}

//解析响应结果
func ParseGetAccountResult(resp []byte) (GetAccountResult, error) {
	getAccountResult := GetAccountResult{}
	err := json.Unmarshal(resp, &getAccountResult)
	if err != nil {
		return getAccountResult, err
	}

	return getAccountResult, nil
}

//往飞语服务器(禁用,启用,修改绑定手机)飞语云账户的响应结果
/**
 * $result:当resultCode为0表示成功
 * 返回的具体说明：
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回样例：
 * {
 *	    "resultCode": "0",
 *	    "resultMsg": "修改账户成功"
 * }
 */
type Result struct {
	ResultCode	string `json:"resultCode"`  //	是	返回的错误代码,0）代表成功；其他具体的错误代码见错误描述
	ResultMsg	string `json:"resultMsg"`  //	否	错误信息描述
}

//解析响应结果
func ParseResult(resp []byte) (Result, error) {
	result := Result{}
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

//查询飞语云账户在线状态的响应结果
/**
 * $result:当resultCode为0表示成功
 * 返回的具体说明：
 * fyAccountId：String必有，飞语云账户ID
 * clientIp：String必有，客户端注册的IP地址
 * onlineTime：String必有，上线时间（UTC 1970-01-01至今的毫秒数）
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回样例：
 * {
	    "result": [
	        {
	            "clientIp": "10.2.3.4",
	            "fyAccountId": "11111",
	            "onlineTime": 1431669626844
	        },
	        {
	            "clientIp": "13.2.3.2",
	            "fyAccountId": "22222",
	            "onlineTime": 1431669626844
	        }
	    ],
	    "resultCode": "0",
	    "resultMsg": "请求成功"
	}
 */
type OnlineStatusResult struct {
	Result []StatusResult `json:"result"` //
	ResultCode	string `json:"resultCode"`  //	是	返回的错误代码,0）代表成功；其他具体的错误代码见错误描述
	ResultMsg	string `json:"resultMsg"`  //	否	错误信息描述
}

type StatusResult struct {
	FyAccountId	string `json:"fyAccountId"` //是	飞语云账户ID
	ClientIp	string `json:"clientIp"` //是	客户端注册的IP地址
	OnlineTime	int64  `json:"onlineTime"` //是	上线时间（UTC 1970-01-01至今的毫秒数）
}

//解析响应结果
func ParseOnlineStatusResult(resp []byte) (OnlineStatusResult, error) {
	onlineStatusResult := OnlineStatusResult{}
	err := json.Unmarshal(resp, &onlineStatusResult)
	if err != nil {
		return onlineStatusResult, err
	}

	return onlineStatusResult, nil
}

//往飞语云服务器主动回拨的接口
/**
 * $result:当resultCode为0表示成功
 * 返回的具体说明：
 * fyCallId：String必有，此次通话呼叫的飞语唯一标识
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回样例：
	 {
		    "result": {
		        "fyCallId": "3343"
		    },
		    "resultCode": "0",
		    "resultMsg": "外呼成功"
		}
 */
type CallBackResult struct {
	FyCallId	string `json:"fyCallId"` //是	此次通话呼叫的飞语唯一标识
	ResultCode	string `json:"resultCode"`  //	是	返回的错误代码,0）代表成功；其他具体的错误代码见错误描述
	ResultMsg	string `json:"resultMsg"`  //	否	错误信息描述
}

//解析响应结果
func ParseCallBackResult(resp []byte) (CallBackResult, error) {
	callBackResult := CallBackResult{}
	err := json.Unmarshal(resp, &callBackResult)
	if err != nil {
		return callBackResult, err
	}

	return callBackResult, nil
}

//飞语的通讯系统会将文本通过语音的形式在用户的手机上播报
/**
 * $result:当resultCode为0表示成功
 * 返回的具体说明：
 * result：(String) 呼叫成功的话，里面存的就是此次呼叫的飞语产生的呼叫ID
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回样例：
	 {
	    "result": "2312erwe33232434@fy",
	    "resultCode": "0",
	    "resultMsg": "外呼成功"
	 }
 */
type VoiceAutoOutCall struct {
	Result	string `json:"result"` //呼叫成功的话，里面存的就是此次呼叫的飞语产生的呼叫ID
	ResultCode	int `json:"resultCode"`  //呼叫状态定义：0）成功；1）失败;
	ResultMsg	string `json:"resultMsg"`  //错误信息描述
}

//解析响应结果
func ParseVoiceAutoOutCall(resp []byte) (VoiceAutoOutCall, error) {
	voiceAutoOutCall := VoiceAutoOutCall{}
	err := json.Unmarshal(resp, &voiceAutoOutCall)
	if err != nil {
		return voiceAutoOutCall, err
	}

	return voiceAutoOutCall, nil
}

/**
 * $result:当resultCode为0表示成功
 * 返回的具体说明：
 * result：(String) 录音文件下载的地址
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回样例：
 {
	 "result": "获取录音文件下载地址",
	 "resultCode": "0",
	 "resultMsg": "外呼成功"
 }
 */
type RecordDownUrl struct {
	Result	string `json:"result"` //呼叫成功的话，里面存的就是此次呼叫的飞语产生的呼叫ID
	ResultCode	string `json:"resultCode"`  //呼叫状态定义：0）成功；1）失败;
	ResultMsg	string `json:"resultMsg"`  //错误信息描述
}

//解析响应结果
func ParseRecordDownUrl(resp []byte) (RecordDownUrl, error) {
	recordDownUrl := RecordDownUrl{}
	err := json.Unmarshal(resp, &recordDownUrl)
	if err != nil {
		return recordDownUrl, err
	}

	return recordDownUrl, nil
}

/**
 * $result:当resultCode为0表示成功
 * 返回的具体说明：
 * result：
 * resultCode：返回的错误代码。0）代表成功；其他具体的错误代码见错误描述
 * resultMsg：错误信息描述
 * 返回样例：
 * {
	    "result": {
	        "appId": "XXXXX",
	        "appCallId": "XXX",
	        "fyCallId": "XXXX",
	        "appServerExtraData": "XXXX",
		 "callbackFirstStartTime": 144454555,
	        "callbackFirstEndTime":144454555,
	        "callStartTime": 144454555,
	        "callEndTime": 144454555
		 "stopReason": 11,
	        "trueShowNumberType": 1,
	        "trueIfRecord": 1
	    },
	    "resultCode": "0",
	    "resultMsg": "获取成功"
	}
 */
type CallHistory struct {
	ResultCode	string `json:"resultCode"` //	接收结果：0）成功；其他代表错误码
	ResultMsg	string `json:"resultMsg"` //	结果描述
	FyCallId	string `json:"fyCallId"` //	飞语产生的呼叫ID
	AppCallId	string `json:"appCallId"` //	第三方呼叫的唯一标识
	AppId	string `json:"appId"` //	应用ID
	AppServerExtraData	string `json:"appServerExtraData"` //	APP附加的数据，通话鉴权的时候产生的
	CallbackFirstStartTime	string `json:"callbackFirstStartTime"` //	回拨第一路开始通话时间，如果是直拨模式，这一路时间是0，回拨才有这一路时间格式为离1970年的毫秒数
	CallbackFirstEndTime	string `json:"callbackFirstEndTime"` //	回拨第一路结束通话时间
	CallStartTime	string `json:"callStartTime"` //	通话的开始时间、回拨第二路的开始时间或者直拨的开始时间
	CallEndTime	string `json:"callEndTime"` //	通话的结束时间、回拨第二路结束通话时间或者直拨的结束时间
	StopReason	string `json:"stopReason"` //	中断原因
	TrueShowNumberType	string `json:"trueShowNumberType"` //	真实的显号类型：1为显示号码，2为不显示号码
	TrueIfRecord	string `json:"trueIfRecord"` //	真实的是否录音：1为录音，0或者2为不录音
}

//解析响应结果
func ParseCallHistory(resp []byte) (CallHistory, error) {
	result := CallHistory{}
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// conference 会议模式

/** 客户端创建会议室，飞语云向第三方服务器授权接口
 * action	String	是	请求类型：创建会议传createConference，加入会议传joinConference
 * callType	Int	是	呼叫类型：5）电话会议
 * fyAccountId	String	是	飞语云账户
 * conferenceId	String	是	会议Id
 * fyCallId	String	是	飞语产生的呼叫的唯一标识
 * appAccountId	String	是	在用户服务器端用户的唯一名称
 * appId	String	是	应用id
 * appExtraData	String	否	第三方私有数据，
 * channelId	String	否	渠道号，做为用户自己统计用
 * ifRecord	Int	是	是否需要录音：1）需要；2）不需要，目前会议不支持录音，字段预留
 * ti	Long	是	自1970年1月1日0时起的毫秒数（UTC/GMT+08:00）
 * au	String	是	MD5（appId+ appToken+ti）
 */
type CallAuthResult struct {
	Action	string `json:"action"` //	是	请求类型：创建会议传createConference，加入会议传joinConference
	CallType	int `json:"callType"` //	是	呼叫类型：5）电话会议
	FyAccountId	string `json:"fyAccountId"` //	是	飞语云账户
	ConferenceId	string `json:"conferenceId"` //	是	会议Id
	FyCallId	string `json:"fyCallId"` //	是	飞语产生的呼叫的唯一标识
	AppAccountId	string `json:"appAccountId"` //	是	在用户服务器端用户的唯一名称
	AppId	string `json:"appId"` //	是	应用id
	AppExtraData	string `json:"appExtraData"` //	否	第三方私有数据，
	ChannelId	string `json:"channelId"` //	否	渠道号，做为用户自己统计用
	IfRecord	int	`json:"ifRecord"` //是	是否需要录音：1）需要；2）不需要，目前会议不支持录音，字段预留
	Ti	int64 `json:"ti"` //是	自1970年1月1日0时起的毫秒数（UTC/GMT+08:00）
	Au	string `json:"au"` //	是	MD5（appId+ appToken+ti）
}

//解析响应结果
func ParseCallAuth(resp []byte) (CallAuthResult, error) {
	result := CallAuthResult{}
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

//解析响应结果
func ParseCallAuthStr(resp []byte) (CallAuthResult, error) {
	result := CallAuthResult{}
	var str []string = strings.Split(string(resp), "&")
	for _, v := range str {
		var val []string = strings.Split(v, "=")
		switch val[0] {
			case "action":
				result.Action = val[1]
			case "appAccountId":
				result.AppAccountId = val[1]
			case "appExtraData":
				result.AppExtraData = val[1]
			case "appId":
				result.AppId = val[1]
			case "callType":
				i, _ := strconv.Atoi(val[1])
				result.CallType = i
			case "channelId":
				result.ChannelId = val[1]
			case "conferenceId":
				result.ConferenceId = val[1]
			case "fyAccountId":
				result.FyAccountId = val[1]
			case "fyCallId":
				result.FyCallId = val[1]
			case "ifRecord":
				i, _ := strconv.Atoi(val[1])
				result.IfRecord = i
			case "ti":
				i64, _ := strconv.ParseInt(val[1], 10, 0)
				result.Ti = i64
			case "au":
				result.Au = val[1]
		}
	}

	return result, nil
}

//响应请求
type CallAuthResponse struct {
	ResultCode	string `json:"resultCode"` //	是	认证授权结果：0）：成功，其他代表错误码（如果自定义的错误码请使用开头9的6位数字，此错误码会透传给客户端）
	ResultMsg	string `json:"resultMsg"` //	否	认证结果描述，会透传给客户端认证授权成功的话，result是下面的json数组
	Result CallAuthResp `json:"result"` //结束
}

type CallAuthResp struct {
	AppCallId	string `json:"appCallId"` //	否	APP端自己产生的呼叫id
	MaxCallMinute int `json:"maxCallMinute"` //	是	最大通话分钟数
	AppServerExtraData string `json:"appServerExtraData"` //	否	APP附加的数据，通话鉴权的时候产生的，话单推送的时候会同时推送APP服务器端
	IfRecord	int `json:"ifRecord"` //	是	是否录音：1为录音，2为不录音
	MaxMember	int `json:"maxMember"` //	否		会议室成员个数，默认系统支持的最大成员数
}

//响应结果
func EncodeCallAuth(resp CallAuthResponse) ([]byte, error) {
	b, err := json.Marshal(&resp)
	return b, err
}

//其中stopReason的返回参数是表示中断原因：
//1:主叫挂断
//2:被叫挂断
//3:呼叫不可及
//5:超时未接
//6:拒接或超时
//7:网络问题
//9:API请求挂断
//10:余额不足
//11:呼叫失败，系统错误
//12:被叫拒接
//13:被叫无人接听
//14:被叫正忙
//15:被叫不在线
//16:呼叫超过最大呼叫时间
/** 当客户端完成呼叫的时候，飞语云服务器会向第三方应用服务器推送此次通话的具体信息
 * action	String	是	请求类型固定值：callhangup
 * appId	String	是	应用ID
 * appCallId	String	否	第三方呼叫的唯一标识
 * fyCallId	String	是	飞语产生的呼叫ID
 * appServerExtraData	String	否	APP附加的数据，通话鉴权的时候产生的
 * callbackFirstStartTime	long	否	回拨第一路开始通话时间，如果是直拨模式，这一路时间是0，回拨才有这一路
 * 时间格式为离1970年的毫秒数
 * callbackFirstEndTime	long	否	回拨第一路结束通话时间
 * callStartTime	long	是	通话的开始时间、回拨第二路的开始时间或者直拨的开始时间
 * callEndTime	long	是	通话的结束时间、回拨第二路结束通话时间或者直拨的结束时间
 * stopReason	int	是	中断原因
 * trueShowNumberType	int	是	真实的显号类型：1为显示号码，2为不显示号码
 * trueIfRecord	int	是	真实的是否录音：1为录音，0或者2为不录音
 * ti	long	是	自1970年1月1日0时起的毫秒数（UTC/GMT+08:00）
 * au	String	是	MD5（appId+ appToken+ti）
 */
type TelListPushResult struct {
    Action	string `json:"action"` //	是	请求类型固定值：callhangup
    AppId	string `json:"appId"` //	是	应用ID
    AppCallId	string `json:"appCallId"` //	否	第三方呼叫的唯一标识
    FyCallId	string `json:"fyCallId"` //	是	飞语产生的呼叫ID
    AppServerExtraData	string `json:"appServerExtraData"` //	否	APP附加的数据，通话鉴权的时候产生的
    CallbackFirstStartTime	int64 `json:"callbackFirstStartTime"` //	否	回拨第一路开始通话时间，如果是直拨模式，这一路时间是0，回拨才有这一路 时间格式为离1970年的毫秒数
    CallbackFirstEndTime	int64 `json:"callbackFirstEndTime"` //	否	回拨第一路结束通话时间
    CallStartTime	int64 `json:"callStartTime"` //	是	通话的开始时间、回拨第二路的开始时间或者直拨的开始时间
    CallEndTime	int64 `json:"callEndTime"` //	是	通话的结束时间、回拨第二路结束通话时间或者直拨的结束时间
    StopReason	int `json:"stopReason"` //	是	中断原因
    TrueShowNumberType	int `json:"trueShowNumberType"` //	是	真实的显号类型：1为显示号码，2为不显示号码
    TrueIfRecord	int `json:"trueIfRecord"` //	是	真实的是否录音：1为录音，0或者2为不录音
    Ti	int64 `json:"ti"` //	是	自1970年1月1日0时起的毫秒数（UTC/GMT+08:00）
    Au	string `json:"au"` //	是	MD5（appId+ appToken+ti）
}

//解析响应结果
func ParseTelListPush(resp []byte) (TelListPushResult, error) {
	result := TelListPushResult{}
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

//解析响应结果
func ParseTelListPushStr(resp []byte) (TelListPushResult, error) {
	result := TelListPushResult{}
	var str []string = strings.Split(string(resp), "&")
	for _, v := range str {
		var val []string = strings.Split(v, "=")
		switch val[0] {
		case "action":
			result.Action = val[1]
		case "appCallId":
			result.AppCallId = val[1]
		case "appServerExtraData":
			result.AppServerExtraData = val[1]
		case "appId":
			result.AppId = val[1]
		case "callbackFirstStartTime":
			i64, _ := strconv.ParseInt(val[1], 10, 0)
			result.CallbackFirstStartTime = i64
		case "callbackFirstEndTime":
			i64, _ := strconv.ParseInt(val[1], 10, 0)
			result.CallbackFirstEndTime = i64
		case "callStartTime":
			i64, _ := strconv.ParseInt(val[1], 10, 0)
			result.CallStartTime = i64
		case "callEndTime":
			i64, _ := strconv.ParseInt(val[1], 10, 0)
			result.CallEndTime = i64
		case "stopReason":
			i, _ := strconv.Atoi(val[1])
			result.StopReason = i
		case "trueShowNumberType":
			i, _ := strconv.Atoi(val[1])
			result.TrueShowNumberType = i
		case "trueIfRecord":
			i, _ := strconv.Atoi(val[1])
			result.TrueIfRecord = i
		case "fyCallId":
			result.FyCallId = val[1]
		case "ti":
			i64, _ := strconv.ParseInt(val[1], 10, 0)
			result.Ti = i64
		case "au":
			result.Au = val[1]
		}
	}

	return result, nil
}

//响应请求
type TelListPushResponse struct {
	ResultCode	string `json:"resultCode"` //	是	认证授权结果：0）：成功，其他代表错误码（如果自定义的错误码请使用开头9的6位数字，此错误码会透传给客户端）
	ResultMsg	string `json:"resultMsg"` //	否	认证结果描述，会透传给客户端认证授权成功的话，result是下面的json数组
}

//响应结果
func EncodeTelListPush(resp TelListPushResponse) ([]byte, error) {
	b, err := json.Marshal(&resp)
	return b, err
}
