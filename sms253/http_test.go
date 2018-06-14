package sms253

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	code := "111111"
	msg := fmt.Sprintf("【253云通讯】你的验证码%s", code)
	phone := "13711111111"
	account := "CN1111111"
	password := "xxxxxxxxxxxxxx"
	targetURL := "http://smssh1.253.com/msg/send/json"
	err := SendSms(targetURL, account, password, phone, msg)
	if err != nil {
		t.Logf("err %v", err)
		t.Logf("send sms failed phone %s, code %s", phone, code)
	} else {
		t.Logf("send sms successful phone %s, code %s", phone, code)
	}
}
