package apple

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func PayVeriy(receipt string) (*TradeResult, error) {
	body, err := ToJson(receipt)
	if err != nil {
		return nil, err
	}
	url := APPLE_VERIY_SANDBOX_URL
	resp, err := doHttpPost(url, body)
	if err != nil {
		return nil, err
	}
	result, err := ParseTradeResult(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// doRequest post the order in json format with a sign
func doHttpPost(targetUrl string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/json")

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		Dial:               TimeoutDialer(5, 5),
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

func TimeoutDialer(dailTimeout, rwTimeout int64) func(network, addr string) (net.Conn, error) {
	return func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, time.Duration(dailTimeout)*time.Second)
		if nil != conn {
			conn.SetDeadline(time.Now().Add(time.Duration(rwTimeout) * time.Second))
		}
		return conn, err
	}
}
