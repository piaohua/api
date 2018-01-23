package feiyu

import (
	"bytes"
	"fmt"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/golang/glog"
)

// FYAPI is abstact of feiyu api handler.
type FYAPI struct {
	Config *FeiyuConfig
}

// Initialized the AppTrans with specific config
func NewFeiyuApi(cfg *FeiyuConfig) (*FYAPI, error) {
	if cfg.APP_ID == "" ||
		cfg.APP_TOKEN == "" ||
		cfg.TEL_LIST_PUSH == "" ||
		cfg.CALL_AUTH == "" {
		return &FYAPI{Config: cfg}, errors.New("config field canot empty string")
	}

	return &FYAPI{Config: cfg}, nil
}

// doRequest get the order in json format with a sign
func doHttpPost(targetUrl string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		ResponseHeaderTimeout: time.Second * 5,
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

/**
 * 添加飞语云账号
 * appAccountId(String)必填，在应用服务器端用户的唯一名称，同一个应用必须要保证唯一
 * globalMobilePhone(String)选填，绑定手机号码。拨打出去可以显示用户的本机号码,要带国别码，如果是中国的是86133*******。如果账户要调用双向回拨接口，必须绑定手机号
 * district(String)，号码的国际区号（中国就是86）
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) AddAccount(accId, phone string) (AddAccountResult, error) {
	addAccountResult := AddAccountResult{}

	queryUrl, postFields := this.newAddAccountParam(accId, phone)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return addAccountResult, err
	}

	addAccountResult, err = ParseAddAccountResult(resp)
	if err != nil {
		return addAccountResult, err
	}

	return addAccountResult, nil
}

//添加飞语云账户请求url
func (this *FYAPI) newAddAccountParam(accId, phone string) (string, string) {
	param := make(map[string]string)
	param["appAccountId"] = accId
	if phone != "" {
		param["globalMobilePhone"] = "86"+phone
		param["district"] = "86"
	}
	return this.Config.ADD_ACCOUNT_URL, this.queryArgs(param)
}

/**
 * 查看终端SDK账户
 * infoType(String)必填，查询信息类型。1）飞语云账户号码；2）APP账户号码；3）手机号码
 * info(String)必填，infoType对应的查询信息，例：infoType=3，info=15xxxxxx(手机号码)
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) GetAccount(infoType, info string) (GetAccountResult, error) {
	getAccountResult := GetAccountResult{}

	queryUrl, postFields := this.newGetAccountParam(infoType, info)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return getAccountResult, err
	}

	getAccountResult, err = ParseGetAccountResult(resp)
	if err != nil {
		return getAccountResult, err
	}

	return getAccountResult, nil
}

//查询飞语云账户请求url
func (this *FYAPI) newGetAccountParam(infoType, info string) (string, string) {
	param := make(map[string]string)
	param["infoType"] = infoType
	param["info"] = info
	return this.Config.GET_ACCOUNT_URL, this.queryArgs(param)
}

/**
 * 往飞语云服务器禁用飞语云账户的接口
 * fyAccountId(String),飞语云账户ID
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) DisableAccount(accId string) bool {

	queryUrl, postFields := this.newDisableAccountParam(accId)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return false
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		return false
	}

	return result.ResultCode == "0"
}

//禁用飞语云账户请求url
func (this *FYAPI) newDisableAccountParam(accId string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	return this.Config.DISABLE_ACCOUNT_URL, this.queryArgs(param)
}

/**
 * 往飞语云服务器启用飞语云账户的接口
 * fyAccountId(String),飞语云账户ID
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) EnableAccount(accId string) bool {

	queryUrl, postFields := this.newEnableAccountParam(accId)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return false
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		return false
	}

	return result.ResultCode == "0"
}

//启用飞语云账户请求url
func (this *FYAPI) newEnableAccountParam(accId string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	return this.Config.ENABLE_ACCOUNT_URL, this.queryArgs(param)
}

/**
 * 修改飞语云账户绑定手机号
 * fyAccountId(String),飞语云账户ID
 * globalMobilePhone(String)必填，待绑定的手机号码。用户显示号码和回拨用,要带国别码;例如：86+13888888888；当手机号为空时候，代表是解除手机号码的绑定
 * district(String)，号码的国际区号（中国就是86）
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) ModifyAccountDisplayNumber(accId, phone string) bool {

	queryUrl, postFields := this.newModifyParam(accId, phone)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return false
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		return false
	}

	return result.ResultCode == "0"
}

//修改飞语云账户绑定手机请求url
func (this *FYAPI) newModifyParam(accId, phone string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	param["globalMobilePhone"] = "86"+phone
	param["district"] = "86"
	return this.Config.MODIFY_ACCOUNT_DISPLAY_NUMBER_URL, this.queryArgs(param)
}

/**
 * 查询飞语云账户的在线状态
 * fyAccountId(String),飞语云账户ID
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数
 */
func (this *FYAPI) OnlineStatus(accId string) (OnlineStatusResult, error) {
	result := OnlineStatusResult{}

	queryUrl, postFields := this.newOnlineParam(accId)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return result, err
	}

	result, err = ParseOnlineStatusResult(resp)
	if err != nil {
		return result, err
	}

	return result, nil
}

//查询飞语云账户在线状态请求url
func (this *FYAPI) newOnlineParam(accId string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	return this.Config.GET_ACCOUNT_ONLINE_STATUS_URL, this.queryArgs(param)
}

/**
 * 集成方服务器端主动发起回
 * caller(String)，主叫号码：（可以填写手机号码，或者飞语云ID，如果是飞语云ID，则此ID必须绑定手机号码）
 * maxCallMinute，此次最大通话分钟数，最大120分钟
 * showNumberType，外呼显号标示：1）显号； 2）不显号
 * callee，被叫号码：号码格式如下，拨打中国手机86+13888888888，拨打中国上海固话86+21+12341234
 * calleeDistrictCode，被叫默认的国别码，默认是86
 * ifRecord(int)，是否需要录音：1）需要；2）不需要
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) CallBack(caller, callId, callee string) (CallBackResult, error) {
	result := CallBackResult{}

	queryUrl, postFields := this.newCallBackParam(caller, callId, callee)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return result, err
	}

	result, err = ParseCallBackResult(resp)
	if err != nil {
		return result, err
	}

	return result, nil
}

//往飞语云服务器主动回拨的接口请求url
func (this *FYAPI) newCallBackParam(caller, callId, callee string) (string, string) {
	param := make(map[string]string)
	param["caller"] = caller
	param["appCallId"] = callId
	param["maxCallMinute"] = "120"
	param["showNumberType"] = "1"
	param["callee"] = callee
	param["calleeDistrictCode"] = "86"
	param["ifRecord"] = "1"
	return this.Config.CALL_BACK_URL, this.queryArgs(param)
}

/**
 * 第三方应用服务器发送相关信息给飞语云服务器。
 * 符合条件的情况下，飞语的通讯系统会将文本通过语音的形式在用户的手机上播报。
 * phone(String)必填，被叫手机号码
 * msg(String)必填，需要拨报的文字内容
 * playtimes(int)必填，播放次数，默认是1
 * replaytimes(int)必填，重听次数，默认是0
 * appCallId(String)必填，第三方呼叫的唯一标识
 * retrytimes(int)必填，呼叫次数，默认是1次就是代表呼叫一次，如果2的话，表示第一次呼叫失败了，会再呼叫一次
 * retrystep(int)必填，重试的时间间隔（秒），默认0秒
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) SoundSmsAction(callId, callee, msg string) (VoiceAutoOutCall, error) {
	result := VoiceAutoOutCall{}

	queryUrl, postFields := this.newSoundSmsActionParam(callId, callee, msg)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return result, err
	}

	result, err = ParseVoiceAutoOutCall(resp)
	if err != nil {
		return result, err
	}

	return result, nil
}

//往飞语云服务器主动回拨的接口请求url
func (this *FYAPI) newSoundSmsActionParam(callId, callee, msg string) (string, string) {
	param := make(map[string]string)
	param["playtimes"] = "2"
	param["replaytimes"] = "1"
	param["appCallId"] = callId
	param["callee"] = callee
	param["msg"] = msg
	return this.Config.CALL_BACK_URL, this.queryArgs(param)
}

/**
 * 获取录音文件下载地址，请在通话完成15分钟后调用此接口获取地址。
 * 获取的下载地址有效期30分钟，30分钟后需要重新请求接口获取新下载地址。
 * fyCallId(String)必填，飞语产生的呼叫ID
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) GetRecordDownUrl(fyCallId string) (RecordDownUrl, error) {
	result := RecordDownUrl{}

	queryUrl, postFields := this.newRecordDownUrlParam(fyCallId)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return result, err
	}

	result, err = ParseRecordDownUrl(resp)
	if err != nil {
		return result, err
	}

	return result, nil
}

//获取录音文件下载地址请求url
func (this *FYAPI) newRecordDownUrlParam(fyCallId string) (string, string) {
	param := make(map[string]string)
	param["fyCallId"] = fyCallId
	return this.Config.CALL_BACK_URL, this.queryArgs(param)
}

/**
 * 电话结束后，可以通过这个接口查询通话记录
 * fyCallId(String)必填，飞语产生的呼叫ID
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) FetchCallHistory(fyCallId string) (CallHistory, error) {
	result := CallHistory{}

	queryUrl, postFields := this.newCallHistoryParam(fyCallId)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return result, err
	}

	result, err = ParseCallHistory(resp)
	if err != nil {
		return result, err
	}

	return result, nil
}

//获取录音文件下载地址请求url
func (this *FYAPI) newCallHistoryParam(fyCallId string) (string, string) {
	param := make(map[string]string)
	param["fyCallId"] = fyCallId
	return this.Config.FETCH_CALL_HISTORY, this.queryArgs(param)
}

/**
 *拼接请求参数
 */
func (this *FYAPI) queryArgs(param map[string]string) string {
	param["appId"] = this.Config.APP_ID //每个请求都有该参数
	//var ti string = TimestampNano()
	var ti string = Timestamp() + "000"
	param["ti"] = ti
	param["ver"] = this.Config.VERSION
	sign := Sign(this.Config.APP_ID, this.Config.APP_TOKEN, ti)
	param["au"] = sign
	return ToQueryString(param)
}

// conference 会议模式

/**
 * 此接口用于离开会议以及要取消整个会议室的时候调用
 * fyAccountId	String	否	飞语云账户ID
 * conferenceId	String	是	飞语云往第三方服务器授权的时候提供的会议号
 * hangupType	int	是	挂断类型：0）挂断当前参数提供的飞语号，1）挂断除当前参数提供的飞语号的其他人，2）挂断所有人，此时飞语号可不传
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) HangUp(accId, cId, hType string) string {

	queryUrl, postFields := this.newHangUpParam(accId, cId, hType)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		glog.Infof("HangUp -> %s, %s, %s, %v", accId, cId, hType, err)
		glog.Infof("HangUp resp -> %v, %v", resp, err)
		return fmt.Sprintf("HangUp err:%v", err)
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		glog.Infof("HangUp result -> %s, %s, %s, %v", accId, cId, hType, err)
		return fmt.Sprintf("HangUp err:%v", err)
	}

	return result.ResultCode
}

//挂断请求url
func (this *FYAPI) newHangUpParam(accId, cId, hType string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	param["conferenceId"] = cId
	param["hangupType"] = hType
	return this.Config.CONFERENCE_HANGUP, this.queryArgs(param)
}

//(离开房间)主动挂断
func HangUpSrv(accId, cId, hType string) string {
	cfg := NewFeiyuConfig()
	fy, err := NewFeiyuApi(cfg)
	if err != nil {
		glog.Infof("HangUp HangUpSrv -> %s, %s, %s, %v", accId, cId, hType, err)
		return fmt.Sprintf("HangUpSrv err:%v", err)
	}
	return fy.HangUp(accId, cId, hType)
}

/**
 * 此接口用于将用户设置静音
 * fyAccountId	String	否	飞语云账户ID
 * conferenceId	String	是	飞语云往第三方服务器授权的时候提供的会议号
 * muteType	int	是	静音类型：0）将当前参数提供的飞语号静音，1）除了当前参数提供的飞语号，其他人设置静音，2）将所有人设置静音，此时飞语号可不传
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) Mute(accId, cId, hType string) string {

	queryUrl, postFields := this.newMuteParam(accId, cId, hType)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return fmt.Sprintf("Mute err:%v", err)
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		return fmt.Sprintf("Mute err:%v", err)
	}

	return result.ResultCode
}

//静音请求url
func (this *FYAPI) newMuteParam(accId, cId, hType string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	param["conferenceId"] = cId
	param["muteType"] = hType
	return this.Config.CONFERENCE_MUTE, this.queryArgs(param)
}

/**
 * 此接口用于将用户设置取消静音
 * fyAccountId	String	否	飞语云账户ID
 * conferenceId	String	是	飞语云往第三方服务器授权的时候提供的会议号
 * unmuteType	int	是	取消静音类型：0）将当前参数提供的飞语号取消静音，1）除了当前参数提供的飞语号，其他人取消静音，2）将所有人取消静音，此时飞语号可不传
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) Unmute(accId, cId, hType string) string {

	queryUrl, postFields := this.newUnmuteParam(accId, cId, hType)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return fmt.Sprintf("Unmute err:%v", err)
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		return fmt.Sprintf("Unmute err:%v", err)
	}

	//return result.ResultCode == "0"
	return result.ResultCode
}

//取消静音请求url
func (this *FYAPI) newUnmuteParam(accId, cId, hType string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	param["conferenceId"] = cId
	param["unmuteType"] = hType
	return this.Config.CONFERENCE_UNMUTE, this.queryArgs(param)
}

/**
 * 此接口用户对会议室进行语音广播
 * fyAccountId	String	否	飞语云账户ID
 * conferenceId	String	是	飞语云往第三方服务器授权的时候提供的会议号
 * playType	int	是	播放类型：0）给当前参数提供的飞语号播放一段语音， 1）将所有播放一段语音
 * fileName	String	是	需要播放的语音文件名称。如：xxx.wav
 * ti(long)必填，时间戳。自1970年1月1日0时起的毫秒数, 时间戳有效时间为30分钟
 */
func (this *FYAPI) Play(accId, cId, fileName, hType string) bool {

	queryUrl, postFields := this.newPlayParam(accId, cId, fileName, hType)
	resp, err := doHttpPost(queryUrl, []byte(postFields))
	if err != nil {
		return false
	}

	result := Result{}
	result, err = ParseResult(resp)
	if err != nil {
		return false
	}

	return result.ResultCode == "0"
}

//播放请求url
func (this *FYAPI) newPlayParam(accId, cId, fileName, hType string) (string, string) {
	param := make(map[string]string)
	param["fyAccountId"] = accId
	param["conferenceId"] = cId
	param["playType"] = hType
	param["fileName"] = fileName
	return this.Config.CONFERENCE_MUTE, this.queryArgs(param)
}
