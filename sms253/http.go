package sms253

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

//SendSms 发送短信
func SendSms(targetURL, account, password, phone, msg string) error {
	bytesData, err := newParam(account, password, phone, msg)
	if err != nil {
		return err
	}
	resp, err := doHTTPPost(targetURL, bytesData)
	if err != nil {
		return err
	}
	//fmt.Printf("resp %s\n", string(resp))
	result, err := parse(resp)
	if err != nil {
		return err
	}
	if result.Code == "0" {
		return nil
	}
	return fmt.Errorf("error code %#v", result)
}

func doHTTPPost(targetURL string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(body))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/json;charset=UTF-8")

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
