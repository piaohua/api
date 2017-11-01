package xfapi

import "encoding/json"

//导入用户
//uid	string	是	导入用户id，唯一标识一个用户
//name	string	是	导入用户的昵称
//props	string	否	用户自定义属性，数据格式可为JSON，也可以是其他数据格式
//icons	string	否	用户头像URL
//cMsgID	string	是	客户端构造，消息标识
type ImportRequest struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Props  string `json:"props"`
	Icons  string `json:"icons"`
	CMsgID string `json:"cMsgID"`
}

func (req *ImportRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//用户信息删除
//uid	string	是	唯一标识一个用户
//cMsgID	string	是	唯一标识一次http请求
type DeleteRequest struct {
	Uid    string `json:"uid"`
	CMsgID string `json:"cMsgID"`
}

func (req *DeleteRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5. 群组功能

//5.1 创建群组
//owner	string	是	群主id（讨论组创建者id）
//gname	string	是	群组（讨论组）名称
//type	int	是	type值为0，表示创建的是讨论组，值为1创建的是群组，其他值将返回错误码
//cMsgID	string	是	消息标识
type GroupCreateRequest struct {
	Owner  string `json:"owner"`
	Gname  string `json:"gname"`
	Type   int    `json:"type"`
	CMsgID string `json:"cMsgID"`
}

func (req *GroupCreateRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.2 添加群组（讨论组）成员
//gid	string	是	群组（讨论组）id
//uid	string	是	发送邀请的用户id
//members	JsonArray	是	添加成员列表
//type	int	是	type为0表示讨论组，为1表示群组，讨论组不需要被拉取人同意，群组需要验证
//msg	string	否	发送邀请时附带的邀请信息
//cMsgID	string	是	消息标识
type GroupAddRequest struct {
	Gid     string   `json:"gid"`
	Uid     string   `json:"uid"`
	Members []string `json:"members"`
	Type    int      `json:"type"`
	Msg     string   `json:"msg"`
	CMsgID  string   `json:"cMsgID"`
}

func (req *GroupAddRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.3 加入群组(仅适用于群组)
//gid	string	是	群组id
//uid	string	是	待加入群组的用户id
//cMsgID	string	是	消息标识
type GroupJoinRequest struct {
	Gid    string `json:"gid"`
	Uid    string `json:"uid"`
	CMsgID string `json:"cMsgID"`
}

func (req *GroupJoinRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.4 退出群组
//gid	string	是	群组id
//uid	string	是	要退群用户的uid
//cMsgID	string	是	消息标识
type GroupExitRequest struct {
	Gid    string `json:"gid"`
	Uid    string `json:"uid"`
	CMsgID string `json:"cMsgID"`
}

func (req *GroupExitRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.5 踢除群组（讨论组）成员
//gid	string	是	群组id
//members	JsonArray	是	待剔除的成员名单
//uid	string	是	调用该接口的用户id，如果是讨论组，只有讨论组创建者才有权限调用该接口，如果选用其他uid则会返回对应的错误码，如果是群组，只有管理员和群主有权限调用该接口，填写其他的uid会返回对应的错误码
//cMsgID	string	是	消息标识
type GroupKickRequest struct {
	Gid     string   `json:"gid"`
	Members []string `json:"members"`
	Uid     string   `json:"uid"`
	CMsgID  string   `json:"cMsgID"`
}

func (req *GroupKickRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.6 解散群组
//gid	string	是	群组id
//uid	string	是	发送该请求的用户，只有群主有权限，其他用户发送会报错
//cMsgID	string	是	消息标识
type GroupRemoveRequest struct {
	Gid    string `json:"gid"`
	Uid    string `json:"uid"`
	CMsgID string `json:"cMsgID"`
}

func (req *GroupRemoveRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.7 获取群组信息
//gid	string	是	群组id
//uid	string	是	用户id
//cMsgID	string	是	消息标识
type GroupInfoRequest struct {
	Gid    string `json:"gid"`
	Uid    string `json:"uid"`
	CMsgID string `json:"cMsgID"`
}

func (req *GroupInfoRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.8 任命管理员
//gid	string	是	群组id
//members	JsonArray	是	待剔除的成员名单
//uid	string	是	调用该接口的用户id，如果是讨论组，只有讨论组创建者才有权限调用该接口，如果选用其他uid则会返回对应的错误码，如果是群组，只有管理员和群主有权限调用该接口，填写其他的uid会返回对应的错误码
//cMsgID	string	是	消息标识
type GroupAppointmgrRequest struct {
	Gid     string   `json:"gid"`
	Members []string `json:"members"`
	Uid     string   `json:"uid"`
	CMsgID  string   `json:"cMsgID"`
}

func (req *GroupAppointmgrRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.9 移除管理员
//gid	string	是	群组id
//members	JsonArray	是	待解除的管理员列表，如{"members":["aaa","bbb"]}
//uid	string	是	调用该接口的用户id，须是群组，其他用户id会返回错误码
//cMsgID	string	是	消息标识
type GroupRevokemgrRequest struct {
	Gid     string   `json:"gid"`
	Members []string `json:"members"`
	Uid     string   `json:"uid"`
	CMsgID  string   `json:"cMsgID"`
}

func (req *GroupRevokemgrRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.10 获取群组列表
//uid	string	是	要查询用户的id
//cMsgID	string	是	消息标识
type GroupListRequest struct {
	Uid    string `json:"uid"`
	CMsgID string `json:"cMsgID"`
}

func (req *GroupListRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.11 群组验证
//type	int	必须	验证类型，1(被邀请人对入群邀请的验证),2(管理员或群主对入群邀请的验证),3(管理员或群主对入群申请的验证)
//oper	int	必须	操作类型 同意（0）拒绝（1）
//uid	string	必须	调用该接口的用户id
//gid	string	必须	群组id
//inviter	string	由type值确定	取值1和2时须携带该字段，邀请人id
//invitee	string	由type值确定	取值1和2时须携带该字段，被邀请人id
//applicant	string	由type值确定	取值3时须携带该字段，申请人id
//cMsgID	string	是	消息标识
type GroupVerifyRequest struct {
	Type      int    `json:"type"`
	Oper      int    `json:"oper"`
	Uid       string `json:"uid"`
	Gid       string `json:"gid"`
	Inviter   string `json:"inviter"`
	Invitee   string `json:"invitee"`
	Applicant string `json:"applicant"`
	CMsgID    string `json:"cMsgID"`
}

func (req *GroupVerifyRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.12 编辑群组资料
//gid	string	是	群组id
//gname	string	否	新的群组名称
//uid	string	是	调用接口的用户id
//announcement	string	否	群公告
//describe	string	否	群描述
//icon	string	否	群头像url
//custom	string	否	自定义属性，数据格式可为JSON，也可以是其他数据格式
//cMsgID	string	是	消息标识
type GroupEditRequest struct {
	Gid          string `json:"gid"`
	Gname        string `json:"gname"`
	Uid          string `json:"uid"`
	Announcement string `json:"announcement"`
	Describe     string `json:"describe"`
	Icon         string `json:"icon"`
	Custom       string `json:"custom"`
	CMsgID       string `json:"cMsgID"`
}

func (req *GroupEditRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.13 移交群主
//gid	string	是	群组id
//uid	string	是	调用接口的用户，须是群主，其他id会报错
//newowner	string	是	新群主id
//cMsgID	string	是	消息标识
type GroupTransferRequest struct {
	Gid      string `json:"gid"`
	Uid      string `json:"uid"`
	Newowner string `json:"newowner"`
	CMsgID   string `json:"cMsgID"`
}

func (req *GroupTransferRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}

//5.14 群组搜索
//condition	string	是	查找条件
//cMsgID	string	是	消息标识
type GroupSearchRequest struct {
	Condition string `json:"condition"`
	CMsgID    string `json:"cMsgID"`
}

func (req *GroupSearchRequest) ToJson() ([]byte, error) {
	return json.Marshal(req)
}
