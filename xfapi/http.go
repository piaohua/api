package xfapi

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Xfapi is abstact of xfyun api handler.
type Xfapi struct {
	Config *XfapiConfig
}

// Initialized the AppTrans with specific config
func NewXfapi(cfg *XfapiConfig) (*Xfapi, error) {
	if cfg.AppId == "" ||
		cfg.AppSecret == "" ||
		cfg.AppToken == "" ||
		cfg.GetTokenUrl == "" ||
		cfg.GetUserTokenUrl == "" ||
		cfg.UpdateUserTokenUrl == "" ||
		cfg.GroupCreateUrl == "" ||
		cfg.GroupAddUrl == "" ||
		cfg.GroupExitUrl == "" {
		return &Xfapi{Config: cfg}, errors.New("config field canot empty string")
	}

	return &Xfapi{Config: cfg}, nil
}

//app级别token获取
//X-Appid	讯飞开放平台注册申请应用的应用ID(appid)	是
//X-Nonce	随机数（最大长度128个字符）	是
//X-CurTime	当前UTC时间戳，从1970年1月1日0点0 分0 秒开始到现在的秒数(String)	是
//X-CheckSum	MD5(apiKey + Nonce + CurTime),三个参数拼接的字符串，进行MD5哈希计算，转化成16进制字符(String，小写)	是
//X-Expiration	token过期时间，单位为妙(s),若希望token不过期，设置为-1，默认1天	否
func (this *Xfapi) GetToken() (*GetTokenResult, error) {

	curtime := timestamp()
	nonce := getMd5String(curtime)
	checksum := generateCheckSum(this.Config.AppSecret, nonce, curtime)
	param := make(map[string]string)
	param["X-Appid"] = this.Config.AppId
	param["X-Nonce"] = nonce
	param["X-CurTime"] = curtime
	param["X-CheckSum"] = checksum
	param["X-Expiration"] = "-1"
	fmt.Println("param ", param)

	resp, err := doHttpGet(this.Config.GetTokenUrl, []byte{}, param)
	fmt.Printf("resp %s, err %v\n", string(resp), err)
	if err != nil {
		return nil, err
	}

	return ParseGetTokenResult(resp)
}

//用户token获取
//X-Appid	讯飞开放平台注册申请应用的应用ID(appid)	是
//X-Token	app级别token，由getToken.do接口获取	是
//X-Uid	用户uid，详见第三部分用户系统对接	是
//X-Expiration	token过期时间，单位为妙(s),若希望token不过期，设置为-1，默认1天	否
func (this *Xfapi) GetUserToken(uid string, update bool) (*GetUserTokenResult, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)
	param["X-Uid"] = uid
	param["X-Expiration"] = "-1"

	targetUrl := this.Config.GetUserTokenUrl
	if update {
		targetUrl = this.Config.UpdateUserTokenUrl
	}

	resp, err := doHttpGet(targetUrl, []byte{}, param)
	if err != nil {
		return nil, err
	}

	return ParseGetUserTokenResult(resp)
}

//导入用户
//request
//uid	string	是	导入用户id，唯一标识一个用户
//name	string	是	导入用户的昵称
//props	string	否	用户自定义属性，数据格式可为JSON，也可以是其他数据格式
//icons	string	否	用户头像URL
//cMsgID	string	是	客户端构造，消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "import userinfo successfully"
//}
func (this *Xfapi) UserImport(uid, name, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(ImportRequest)
	req.Uid = uid
	req.Name = name
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.UserImportUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//用户信息更新
//name、props、icons三个参数不能同时为空值
//request
//uid	string	是	用户id，唯一标识一个用户
//name	string	否	用户新的昵称
//props	string	否	用户自定义属性，数据格式可为JSON，也可以是其他数据格式
//icons	string	否	用户头像URL
//cMsgID	string	是	唯一标识一次http请求
//
//response
//{
//    "ret"       :   0,
//}
func (this *Xfapi) UserUpdate(uid, name, props, icons, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(ImportRequest)
	req.Uid = uid
	req.Name = name
	req.Props = props
	req.Icons = icons
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.UserUpdateUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//用户信息删除
//request
//uid	string	是	唯一标识一个用户
//cMsgID	string	是	唯一标识一次http请求
//
//response
//{
//    "ret"       :   0,
//}
func (this *Xfapi) UserDelete(uid, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(DeleteRequest)
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.UserDeleteUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//TODO
//消息功能
//用户通知
//普通消息rest

//5. 群组功能

//5.1 创建群组
//request
//owner	string	是	群主id（讨论组创建者id）
//gname	string	是	群组（讨论组）名称
//type	int	是	type值为0，表示创建的是讨论组，值为1创建的是群组，其他值将返回错误码
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "gid"       :   "100001",
//    "detail"    :   "create group successfully"
//}
func (this *Xfapi) GroupCreate(owner, gname, cMsgID string, gtype int) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupCreateRequest)
	req.Owner = owner
	req.Gname = gname
	req.Type = gtype
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupCreateUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.2 添加群组（讨论组）成员
//request
//gid	string	是	群组（讨论组）id
//uid	string	是	发送邀请的用户id
//members	JsonArray	是	添加成员列表
//type	int	是	type为0表示讨论组，为1表示群组，讨论组不需要被拉取人同意，群组需要验证
//msg	string	否	发送邀请时附带的邀请信息
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "add members successfully"
//}
func (this *Xfapi) GroupAdd(gid, uid, cMsgID, msg string, members []string, gtype int) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupAddRequest)
	req.Gid = gid
	req.Uid = uid
	req.Members = members
	req.Type = gtype
	req.Msg = msg
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupAddUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.3 加入群组(仅适用于群组)
//request
//gid	string	是	群组id
//uid	string	是	待加入群组的用户id
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "wait mananger accept"
//}
func (this *Xfapi) GroupJoin(gid, uid, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupJoinRequest)
	req.Gid = gid
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupJoinUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.4 退出群组
//request
//gid	string	是	群组id
//uid	string	是	要退群用户的uid
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "exit group successfully"
//}
func (this *Xfapi) GroupExit(gid, uid, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupExitRequest)
	req.Gid = gid
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupExitUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.5 踢除群组（讨论组）成员
//接口描述
//1.讨论组中只有讨论组创建者有权限踢除成员
//2.群组中管理员可以踢除普通成员，但是不能踢除其他管理员，群主拥有最高权限，可踢除所有人
//request
//gid	string	是	群组id
//members	JsonArray	是	待剔除的成员名单
//uid	string	是	调用该接口的用户id，如果是讨论组，只有讨论组创建者才有权限调用该接口，如果选用其他uid则会返回对应的错误码，如果是群组，只有管理员和群主有权限调用该接口，填写其他的uid会返回对应的错误码
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "kick members successfully"
//}
func (this *Xfapi) GroupKick(gid, uid, cMsgID string, members []string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupKickRequest)
	req.Gid = gid
	req.Members = members
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupKickUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.6 解散群组
//request
//gid	string	是	群组id
//uid	string	是	发送该请求的用户，只有群主有权限，其他用户发送会报错
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "remove group successfully"
//}
func (this *Xfapi) GroupRemove(gid, uid, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupRemoveRequest)
	req.Gid = gid
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupRemoveUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.7 获取群组信息
//此接口用户获取群组的基本信息，如群主，群大小，群成员列表等信息
//request
//gid	string	是	群组id
//uid	string	是	用户id
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "get groupinfo successfully",
//    "info"      :   {
//                        "appid"         :   "574faab4"
//                        "gname"         :   "UserDemo",
//                        "gid"           :   "10000",
//                        "members"       :   ["Jack","Tom","Amy"],
//                        "type"          :   0,
//                        "owner"         :   "admin",
//                        "managers"      :   ["Jack","Tom"]
//                        "maxusers"      :   200
//                        "announcement"  :   "IM"
//                        "describe"      :   "this is a test group"
//                    }
//}
func (this *Xfapi) GroupInfo(gid, uid, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupInfoRequest)
	req.Gid = gid
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupInfoUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.8 任命管理员
//用户任命管理员，只有群主有这样的权限，其他用户操作会返回错误码
//request
//gid	string	是	群组id
//members	JsonArray	是	待添加为管理员的名单，如 {"members":["aaa","bbb"]}
//uid	string	是	调用该接口的用户id，id必须是群主，其他用户id会报错
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "Appoint managers successfully"
//}
func (this *Xfapi) GroupAppointmgr(gid, uid, cMsgID string, members []string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupAppointmgrRequest)
	req.Gid = gid
	req.Members = members
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupAppointmgrUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.9 移除管理员
//移除管理员，只有群主有该权限
//request
//gid	string	是	群组id
//members	JsonArray	是	待解除的管理员列表，如{"members":["aaa","bbb"]}
//uid	string	是	调用该接口的用户id，须是群组，其他用户id会返回错误码
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "revoke managers successfully"
//}
func (this *Xfapi) GroupRevokemgr(gid, uid, cMsgID string, members []string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupRevokemgrRequest)
	req.Gid = gid
	req.Members = members
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupRevokemgrUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.10 获取群组列表
//获取一个用户加入的群组及讨论列表，返回基本的群组（讨论组）信息
//request
//uid	string	是	要查询用户的id
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "grouplist" :   [
//                        {"owner":"admin","gname":"test_1","maxusers":200,"gid":"10001","size":6,"type":0},
//                        {"owner":"admin","gname":"test_2","maxusers":200,"gid":"10002","size":2,"type":1},
//                        {"owner":"admin","gname":"test_3","maxusers":200,"gid":"10003","size":2,"type":1}
//                    ]
//}
func (this *Xfapi) GroupList(uid, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupListRequest)
	req.Uid = uid
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupListUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.11 群组验证
//1.该接口是群组验证接口
//2.当邀请人添加被邀请人入群时，需要被邀请人验证（同意或拒绝）;当邀请人不是管理员时，被邀请人同意入群邀请后，需要管理员或群主的验证通过后方可入群；用户主动加群时，需要管理员或是群主的验证通过方可入群
//request
//type	int	必须	验证类型，1(被邀请人对入群邀请的验证),2(管理员或群主对入群邀请的验证),3(管理员或群主对入群申请的验证)
//oper	int	必须	操作类型 同意（0）拒绝（1）
//uid	string	必须	调用该接口的用户id
//gid	string	必须	群组id
//inviter	string	由type值确定	取值1和2时须携带该字段，邀请人id
//invitee	string	由type值确定	取值1和2时须携带该字段，被邀请人id
//applicant	string	由type值确定	取值3时须携带该字段，申请人id
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//}
func (this *Xfapi) GroupVerify(gtype, oper int, uid, gid, inviter, invitee, applicant, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupVerifyRequest)
	req.Type = gtype
	req.Oper = oper
	req.Uid = uid
	req.Gid = gid
	req.Inviter = inviter
	req.Invitee = invitee
	req.Applicant = applicant
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupVerifyUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.12 编辑群组资料
//1.该接口用户编辑群组（讨论组）资料
//2.只有管理员或群主有权限编辑群资料
//request
//gid	string	是	群组id
//gname	string	否	新的群组名称
//uid	string	是	调用接口的用户id
//announcement	string	否	群公告
//describe	string	否	群描述
//icon	string	否	群头像url
//custom	string	否	自定义属性，数据格式可为JSON，也可以是其他数据格式
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "update groupinfo successfully"
//}
func (this *Xfapi) GroupEdit(gid, gname, uid, announcement, describe, icon, custom, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupEditRequest)
	req.Gid = gid
	req.Gname = gname
	req.Uid = uid
	req.Announcement = announcement
	req.Describe = describe
	req.Icon = icon
	req.Custom = custom
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupEditUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.13 移交群主
//该接口用户移交群主给群组中的其他人
//request
//gid	string	是	群组id
//uid	string	是	调用接口的用户，须是群主，其他id会报错
//newowner	string	是	新群主id
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "detail"    :   "transfer owner successfully"
//}
func (this *Xfapi) GroupTransfer(gid, uid, newowner, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupTransferRequest)
	req.Gid = gid
	req.Uid = uid
	req.Newowner = newowner
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupTransferUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

//5.14 群组搜索
//该接口用于群组查找，返回一切匹配或包含查找条件的群组
//request
//condition	string	是	查找条件
//cMsgID	string	是	消息标识
//
//response
//{
//    "ret"       :   0,
//    "grouplist" :   [
//                        {"owner":"admin","gname":"test_1","maxusers":200,"gid":"10001","size":6,"type":0},
//                        {"owner":"admin","gname":"test_2","maxusers":200,"gid":"10002","size":2,"type":1},
//                        {"owner":"admin","gname":"test_3","maxusers":200,"gid":"10003","size":2,"type":1}
//                    ]
//}
func (this *Xfapi) GroupSearch(condition, cMsgID string) (*Result, error) {

	param := newParam(this.Config.AppId, this.Config.AppToken)

	req := new(GroupSearchRequest)
	req.Condition = condition
	req.CMsgID = cMsgID
	body, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	resp, err := doHttpPost(this.Config.GroupSearchUrl, body, param)
	if err != nil {
		return nil, err
	}

	return ParseResult(resp)
}

// doRequest get the order in json format with a sign
func doHttpGet(targetUrl string, body []byte, param map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	for key, val := range param {
		req.Header.Add(key, val)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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

// doRequest post the order in json format with a sign
func doHttpPost(targetUrl string, body []byte, param map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/json;charset=UTF-8")
	for key, val := range param {
		req.Header.Add(key, val)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
