package feiyu

import (
	//"bytes"
	//"fmt"
	//"io/ioutil"
	"net/http"
)

// 使用
// go FYAPI.RecvCallAuth(CallAuth) //goroutine

//接收呼叫授权请求
func (this *FYAPI) RecvCallAuth(recv func(http.ResponseWriter, *http.Request)) {
	http.Handle(this.Config.CALL_AUTH, http.HandlerFunc(recv))
}

/*
// 客户端创建会议室，飞语云向第三方服务器授权接口
func CallAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "text/plain;;charset=UTF-8")
	var buf bytes.Buffer
	resp := CallAuthResponse{}
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		result, err := ParseCallAuth(b)
		fmt.Println(result)
		if err == nil {
			//成功,验证 
			resp.ResultCode = "0"
			resp.ResultMsg = "审核成功"
			resp.Result = CallAuthResp{
				AppCallId: "dddd",
				AppServerExtraData: "test",
				IfRecord: 2,
				MaxCallMinute: 120,
				MaxMember: 4,
			}
		} else {
			//失败
			resp.ResultCode = "900001"
			resp.ResultMsg = "没有收到数据"
		}
	}
	b, err := EncodeCallAuth(resp)
	if err != nil {
		//
	}
	r.Body.Close()
	fmt.Fprintf(&buf, string(b))
	w.Write(buf.Bytes())
}
*/

//验证
func (this *FYAPI) CallAuthVerify(result CallAuthResult) bool {
	if this.Config.APP_ID != result.AppId {
		return false
	}
	sign := Sign(this.Config.APP_ID, this.Config.APP_TOKEN, Time2Str(result.Ti))
	if sign != result.Au {
		return false
	}
	return true
}

// 使用
// go FYAPI.RecvTelListPush(TelListPush) //goroutine

//接收话单推送请求
func (this *FYAPI) RecvTelListPush(recv func(http.ResponseWriter, *http.Request)) {
	http.Handle(this.Config.TEL_LIST_PUSH, http.HandlerFunc(recv))
}

/*
// 客户端创建会议室，飞语云向第三方服务器授权接口
func TelListPush(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "text/plain;;charset=UTF-8")
	var buf bytes.Buffer
	resp := TelListPushResponse{}
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		result, err := ParseTelListPush(b)
		fmt.Println(result)
		if err == nil {
			//成功,验证 
			resp.ResultCode = "0"
			resp.ResultMsg = "接收成功"
		} else {
			//失败
			resp.ResultCode = "900001"
			resp.ResultMsg = "没有收到数据"
		}
	}
	b, err := EncodeTelListPush(resp)
	if err != nil {
		//
	}
	r.Body.Close()
	fmt.Fprintf(&buf, string(b))
	w.Write(buf.Bytes())
}
*/

//验证
func (this *FYAPI) TelListPushVerify(result TelListPushResult) bool {
	if this.Config.APP_ID != result.AppId {
		return false
	}
	sign := Sign(this.Config.APP_ID, this.Config.APP_TOKEN, Time2Str(result.Ti))
	if sign != result.Au {
		return false
	}
	return true
}
