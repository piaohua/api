package smsbao

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	code := "111111"
	msg := fmt.Sprintf("【253云通讯】你的验证码%s", code)
	phone := "13700000000"
	account := "11111111111"
	password := "78ece487e0591c4e17160d5d6ec8bbd4"
	targetURL := ""
	err := SendSms(targetURL, account, password, phone, msg)
	if err != nil {
		t.Logf("err %v", err)
		t.Logf("send sms failed phone %s, code %s", phone, code)
	} else {
		t.Logf("send sms successful phone %s, code %s", phone, code)
	}
}
