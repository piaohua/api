package smsbao

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 发送
func SendSmsbao(phone, content, username, password string) error {

	str := fmt.Sprintf("http://api.smsbao.com/sms?u=%s&p=%s&m=%s&c=%s",
		username, password, phone, urlencoder(content))

	resp, err := doHttpGet(str)
	if err != nil {
		return err
	}
	//fmt.Printf("resp %s\n", string(resp))
	if string(resp) == "0" {
		return nil
	}
	return fmt.Errorf("error code %s", string(resp))
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
