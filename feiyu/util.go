package feiyu

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"time"
	"strconv"
	"unicode"
)

// ToQueryString convert the map[string]string to query string parameters
func ToQueryString(param map[string]string) string {
	var args string
	for k, v := range param {
		if args == "" {
			args += fmt.Sprintf("%s=%s", k, v)
		} else {
			args += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	return args
}

/**
*创建md5摘要,规则是:MD5（appId+ appToken+ti）
*/
func Sign(id, token, ti string) string {
	return ToUpper(Md5(id + token + ti))
}

//golang make the caracter in a string uppercase
func ToUpper(str string) string {
	var s string
	for _, v := range str {
		s += string(unicode.ToUpper(v))
	}
	return s
}

// 获取当前时间截
func Timestamp() string {
	return Time2Str(time.Now().Unix())
}

func Time2Str(ti int64) string {
	return strconv.FormatInt(ti, 10)
}

// 获取当前纳秒时间
func TimestampNano() string {
	return Time2Str(time.Now().UnixNano())
}

// md5 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
