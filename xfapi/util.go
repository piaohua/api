package xfapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// ToQueryString convert the map[string]string to query string parameters
func ToQueryString(param map[string]string) string {
	var args string
	for k, v := range param {
		if args == "" {
			args += fmt.Sprintf("?%s=%s", k, v)
		} else {
			args += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	return args
}

func newParam(AppId, AppToken string) map[string]string {
	param := make(map[string]string)
	param["X-Appid"] = AppId
	param["X-Token"] = AppToken
	return param
}

func generateCheckSum(apikey, nonce, curtime string) string {
	input := fmt.Sprintf("%s%s%s", apikey, nonce, curtime)
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 生成32位md5字符串
 * @param s string
 * @return string
 */
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/**
* 获取当前时间截
* @return string
 */
func timestamp() string {
	return strconv.FormatInt(time.Now().Local().Unix(), 10)
}
