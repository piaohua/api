package xfapi

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var XfApi *Xfapi

func XfApiInit() {
	host := "http://im.voicecloud.cn:1208/v1"
	cfg := &XfapiConfig{
		AppId:              "509bde9a",
		AppSecret:          "b54350b014bd95ace23648d7741c3107",
		AppToken:           "a49170c6-e648-4bfd-8cfc-d2b6155b19d4",
		GetTokenUrl:        host + "/rest/getToken.do",
		GetUserTokenUrl:    host + "/rest/getUserToken.do",
		UpdateUserTokenUrl: host + "/rest/updateUserToken.do",
		GroupCreateUrl:     host + "/group/create.do",
		GroupAddUrl:        host + "/group/add.do",
		GroupExitUrl:       host + "/group/exit.do",
		GroupListUrl:       host + "/group/getlist.do",
		GroupInfoUrl:       host + "/group/getinfo.do",
	}
	xf, err := NewXfapi(cfg)
	if err != nil {
		panic(err)
	}
	XfApi = xf
}

func TestGet(t *testing.T) {
	XfApiInit()
	r, err := XfApi.GetToken()
	t.Logf("%#v", r)
	t.Log(err)
	r1, err1 := XfApi.GetUserToken("100", false)
	t.Logf("%#v", r1)
	t.Log(err1)
}

//curl -XGET -H 'X-Appid: 509bde9a' -H 'X-Nonce: e0722faab1818a4067d8578e011ec0fb' -H 'X-CurTime: 1494061018' -H 'X-CheckSum: f8fc55af8261141720fc4b91f4ce5051' https://im.voicecloud.cn/v1/rest/getToken.do
//param  map[
//X-Appid:509bde9a
//X-Nonce:e0722faab1818a4067d8578e011ec0fb
//X-CurTime:1494061018
//X-CheckSum:f8fc55af8261141720fc4b91f4ce5051
//X-Expiration:-1]

func TestResult(t *testing.T) {
	str := `{"ret":31018,"detail":"service inner error"}`
	//str := `{
	//	"ret"       :   0,
	//	"detail"    :   "Appoint managers successfully"
	//}`
	r, err := ParseResult([]byte(str))
	t.Logf("r %#v", r)
	t.Log(err)
}

func TestInfo(t *testing.T) {
	XfApiInit()
	uid := "11956" //创建者
	//get group list
	cMsgID := bson.NewObjectId().Hex()
	r, err := XfApi.GroupList(uid, cMsgID)
	t.Logf("r %#v, err %v", r, err)
	//get group info
	gid := "af60a112-ca88-4db3-b438-df1ebc305175"
	cMsgID = bson.NewObjectId().Hex()
	r, err = XfApi.GroupInfo(gid, uid, cMsgID)
	t.Logf("r %#v, err %v", r, err)
}

func TestGroup(t *testing.T) {
	XfApiInit()
	uid := "10101"    //创建者
	member := "10106" //加入成员
	rid := "1000"     //讨论组名字
	gid := groupCreate(uid, rid)
	t.Log("gid ", gid)
	//
	if gid != "" {
		ok := groupAdd(uid, gid, member)
		t.Log("add  ", ok)
		if ok {
			ok := groupExit(member, gid)
			t.Log("exit  ", ok)
		}
	}
	//get group list
	cMsgID := bson.NewObjectId().Hex()
	r, err := XfApi.GroupList(uid, cMsgID)
	t.Logf("r %#v, err %v", r, err)
	//if r != nil {
	//	for _, v := range r.Grouplist {
	//		ok := groupExit(uid, v.Gid)
	//		t.Log("exit  ", ok)
	//	}
	//}
	//get group info
	cMsgID = bson.NewObjectId().Hex()
	r, err = XfApi.GroupInfo(gid, uid, cMsgID)
	t.Logf("r %#v, err %v", r, err)
	//gid := "9c3a872d-ff93-48dc-a621-232bf2477112"
	ok := groupExit(uid, gid)
	t.Log("exit  ", ok)
}

//userid 创建者(房主), rid 房间id(讨论组名字)
//return 讨论组id
func groupCreate(userid, rid string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("groupCreate err ", err)
		}
	}()
	cMsgID := bson.NewObjectId().Hex()
	r, err := XfApi.GroupCreate(userid, rid, cMsgID, 0)
	if err != nil || r == nil {
		return ""
	}
	if r.Ret == 0 {
		return r.Gid
	}
	return ""
}

//userid 创建者(房主), gid 讨论组id, member 加入成员id
//return 是否加入成功
func groupAdd(userid, gid, member string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("groupAdd err ", err)
		}
	}()
	cMsgID := bson.NewObjectId().Hex()
	r, err := XfApi.GroupAdd(gid, userid, cMsgID, "", []string{member}, 0)
	if err != nil || r == nil {
		return false
	}
	if r.Ret == 0 {
		return true
	}
	return false
}

//userid 退出者id, gid 讨论组id
//return 是否退出成功
func groupExit(userid, gid string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("groupExit err ", err)
		}
	}()
	cMsgID := bson.NewObjectId().Hex()
	r, err := XfApi.GroupExit(gid, userid, cMsgID)
	if err != nil || r == nil {
		return false
	}
	if r.Ret == 0 {
		return true
	}
	return false
}
