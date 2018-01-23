package wxsdk

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Wxsdk is abstact of weixin sdk handler.
type Wxsdk struct {
	Config *WxsdkConfig
}

// Initialized the AppTrans with specific config
func NewWxsdk(cfg *WxsdkConfig) (*Wxsdk, error) {
	if cfg.AppKey == "" ||
		cfg.CreateRoomUrl == "" ||
		cfg.CancelRoomUrl == "" ||
		cfg.GetCardsUrl == "" ||
		cfg.JsSdkUrl == "" {
		return &Wxsdk{Config: cfg}, errors.New("config field canot empty string")
	}

	return &Wxsdk{Config: cfg}, nil
}

// 创建房间
func (this *Wxsdk) CreateRoom(rid, uuid, num string) (Result, error) {
	result := Result{}

	queryUrl := this.newCreateParam(rid, uuid, num)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return result, err
	}

	result, err = ParseResult(resp)
	if err != nil {
		return result, err
	}

	//verity ErrCode of response
	if result.Code != "1" {
		return result, fmt.Errorf("return Code:%s, Msg:%s", result.Code, result.Msg)
	}

	return result, nil
}

func (this *Wxsdk) newCreateParam(rid, uuid, num string) string {
	param := make(map[string]string)
	param["game"] = "1"
	param["room_id"] = rid
	param["openid"] = uuid
	param["card_num"] = num
	param["time"] = NewTimestampString()

	sign := Sign(param, this.Config.AppKey)
	param["sign"] = sign

	return this.Config.CreateRoomUrl + ToQueryString2(param)
}

// 返还房卡
func (this *Wxsdk) CancelRoom(rid, uuid string) (Result, error) {
	result := Result{}

	queryUrl := this.newCancelParam(rid, uuid)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return result, err
	}

	result, err = ParseResult(resp)
	if err != nil {
		return result, err
	}

	//verity ErrCode of response
	if result.Code != "1" {
		return result, fmt.Errorf("return Code:%s, Msg:%s", result.Code, result.Msg)
	}

	return result, nil
}

func (this *Wxsdk) newCancelParam(rid, uuid string) string {
	param := make(map[string]string)
	param["game"] = "1"
	param["room_id"] = rid
	param["openid"] = uuid
	param["time"] = NewTimestampString()

	sign := Sign(param, this.Config.AppKey)
	param["sign"] = sign

	return this.Config.CancelRoomUrl + ToQueryString2(param)
}

// 查询房卡
func (this *Wxsdk) GetCards(uuid string) (CardsResult, error) {
	result := CardsResult{}

	queryUrl := this.newCardsParam(uuid)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return result, err
	}

	result, err = ParseCardsResult(resp)
	if err != nil {
		return result, err
	}

	//verity ErrCode of response
	if result.Code != "1" {
		return result, fmt.Errorf("return Code:%s, Msg:%s", result.Code, result.Msg)
	}

	return result, nil
}

func (this *Wxsdk) newCardsParam(uuid string) string {
	param := make(map[string]string)
	param["openid"] = uuid
	param["time"] = NewTimestampString()

	sign := Sign(param, this.Config.AppKey)
	param["sign"] = sign

	return this.Config.GetCardsUrl + ToQueryString2(param)
}

// 获取微信JSSDK配置信息
func (this *Wxsdk) JsSdkInfo(url string) (JsSdkResult, error) {
	result := JsSdkResult{}

	queryUrl := this.newInfoParam(url)
	//fmt.Printf("queryUrl %s", queryUrl)
	resp, err := doHttpPost(queryUrl, []byte{})
	if err != nil {
		return result, err
	}

	result, err = ParseJsSdkResult(resp)
	if err != nil {
		return result, err
	}

	//verity ErrCode of response
	if result.Code != "1" {
		return result, fmt.Errorf("return Code:%s, Msg:%s", result.Code, result.Msg)
	}

	return result, nil
}

func (this *Wxsdk) newInfoParam(url string) string {
	param := make(map[string]string)
	param["url"] = url
	param["time"] = NewTimestampString()

	sign := Sign(param, this.Config.AppKey)
	param["sign"] = sign

	return this.Config.JsSdkUrl + ToQueryString2(param)
}

// 验证
func (this *Wxsdk) VerifyLogin(openid, nickname, headimgurl, roomid, sign string, sex, time, vip uint32) error {
	if Timestamp()-int64(time) > 1800 {
		return fmt.Errorf("timeout")
	}
	param := make(map[string]string)
	param["openid"] = openid
	param["nickname"] = nickname
	param["headimgurl"] = headimgurl
	param["room_id"] = roomid
	param["time"] = strconv.FormatInt(int64(time), 10)
	param["sex"] = strconv.FormatInt(int64(sex), 10)
	param["vip"] = strconv.FormatInt(int64(vip), 10)

	if sign != Sign(param, this.Config.AppKey) {
		return fmt.Errorf("sign error")
	}

	return nil
}

// doRequest get the order in json format with a sign
func doHttpPost(targetUrl string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("GET", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: false},
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
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
