package smsbao

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	code := "111111"
	content := fmt.Sprintf("【银行】你的验证码%s，请勿泄露。", code)
	phone := "13700000000"
	username := "11111111111"
	password := "78ece487e0591c4e17160d5d6ec8bbd4"
	err := SendSmsbao(phone, content, username, password)
	if err != nil {
		t.Logf("err %v", err)
		t.Logf("send sms failed phone %s, code %s", phone, code)
	} else {
		t.Logf("send sms successful phone %s, code %s", phone, code)
	}
}
