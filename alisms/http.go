package alisms

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 发送
func SendSms(phone, SignName, TemplateCode, TemplateParam,
	AccessKeyId, AccessKeySecret string) {
	m := newParam(phone, SignName, TemplateCode,
		TemplateParam, AccessKeyId)
	str := SortAndConcat(m)
	str = stringToSign(str)
	str = sign(str, AccessKeySecret)
	str = specialUrlEncode(str)

	str = "http://dysmsapi.aliyuncs.com/?" + str

	resp, err := doHttpGet(str)
	if err != nil {
		fmt.Printf("err %v\n", err)
	}
	fmt.Printf("resp %s\n", string(resp))
}

func doHttpGet(targetUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetUrl, bytes.NewBuffer([]byte{}))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "text/plain;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
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
