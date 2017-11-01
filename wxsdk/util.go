package wxsdk

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
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

// ToQueryString2 convert the map[string]string to query string parameters
func ToQueryString2(param map[string]string) string {
	var args string
	for k, v := range param {
		args += fmt.Sprintf("&%s=%s", k, v)
	}

	return args
}

// SortAndConcat sort the map by key in ASCII order,
// and concat it in form of "k1=v1&k2=2"
func SortAndConcat(param map[string]string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}

	var sortedParam []string
	sort.Strings(keys)
	for _, k := range keys {
		// fmt.Println(k, "=", param[k])
		sortedParam = append(sortedParam, k+"="+param[k])
	}

	return strings.Join(sortedParam, "&")
}

// Sign the parameter in form of map[string]string with app key.
// Empty string and "sign" key is excluded before sign.
// Please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=4_3
func Sign(param map[string]string, key string) string {
	newMap := make(map[string]string)
	// fmt.Printf("%#v\n", param)
	for k, v := range param {
		if k == "sign" {
			continue
		}
		if v == "" {
			continue
		}
		newMap[k] = v
	}
	// fmt.Printf("%#v\n\n", newMap)

	preSignStr := SortAndConcat(newMap)
	preSignWithKey := preSignStr + "&key=" + key

	return fmt.Sprintf("%x", md5.Sum([]byte(preSignWithKey)))
}

const ChinaTimeZoneOffset = 8 * 60 * 60 //Beijing(UTC+8:00)

// NewTimestampString return
func NewTimestampString() string {
	//return fmt.Sprintf("%d", time.Now().Unix()+ChinaTimeZoneOffset)
	return fmt.Sprintf("%d", time.Now().Unix())
}

func Timestamp() int64 {
	return time.Now().Unix()
}
