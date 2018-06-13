package alisms

import "testing"

func TestSend(t *testing.T) {
	phone := "13711111111"
	SignName := "云通信"
	TemplateCode := "SMS_1000000"
	TemplateParam := "{\"name\":\"Tom\", \"code\":\"123\"}"
	AccessKeyId := "11111111111"
	AccessKeySecret := "78ece487e0591c4e17160d5d6ec8bbd4"
	SendSms(phone, SignName, TemplateCode, TemplateParam,
		AccessKeyId, AccessKeySecret)
}
