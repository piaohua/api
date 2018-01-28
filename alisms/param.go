package alisms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

/* 请求参数
AccessKeyId	是
Timestamp	是	格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT
Format	否	没传默认为JSON，可选填值：XML
SignatureMethod	是	建议固定值：HMAC-SHA1
SignatureVersion	是	建议固定值：1.0
SignatureNonce	是	用于请求的防重放攻击，每次请求唯一，JAVA语言建议用：java.util.UUID.randomUUID()生成即可
Signature	是	最终生成的签名结果值

Action	是	API的命名，固定值，如发送短信API的值为：SendSms
Version	是	API的版本，固定值，如短信API的值为：2017-05-25
RegionId	是	API支持的RegionID，如短信API的值为：cn-hangzhou
PhoneNumbers	是	具体见API文档描述
SignName	是	具体见API文档描述
TemplateCode	是	具体见API文档描述
TemplateParam	否	具体见API文档描述
OutId	否	具体见API文档描述
*/

func newParam(phone, SignName, TemplateCode, TemplateParam,
	AccessKeyId string) map[string]string {
	sms := make(map[string]string)
	//系统参数
	sms["AccessKeyId"] = AccessKeyId     //AccessKeyId
	sms["Timestamp"] = timestamp()       //格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT
	sms["Format"] = "json"               //返回数据类型,支持xml,json
	sms["SignatureMethod"] = "HMAC-SHA1" //固定参数
	sms["SignatureVersion"] = "1.0"      //固定参数
	sms["SignatureNonce"] = uniqid()     //用于请求的防重放攻击，每次请求唯一
	//sms["Signature"] = sign(str, accessKeySecret)  //最终生成的签名结果值
	//业务参数
	sms["Action"] = "SendSms"            //api命名 固定值
	sms["Version"] = "2017-05-25"        //api版本 固定值
	sms["RegionId"] = "cn-hangzhou"      //固定参数
	sms["PhoneNumbers"] = phone          //手机号
	sms["SignName"] = SignName           //签名
	sms["TemplateCode"] = TemplateCode   //短信模版id
	sms["TemplateParam"] = TemplateParam //模版内容
	return sms
}

func timestamp() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}

func uniqid() string {
	return uuid.New().String()
}

func urlencoder(str string) string {
	return url.QueryEscape(str)
}

func SortAndConcat(param map[string]string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}

	var sortedParam []string
	sort.Strings(keys)
	for _, k := range keys {
		sortedParam = append(sortedParam, specialUrlEncode(k)+"="+specialUrlEncode(param[k]))
	}

	return strings.Join(sortedParam, "&")
}

func specialUrlEncode(str string) string {
	str = urlencoder(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

//HTTPMethod + “&” + specialUrlEncode(“/”) + ”&” + specialUrlEncode(sortedQueryString)
func stringToSign(str string) string {
	return "GET&%2F" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(str)
}

func sign(str, accessKeySecret string) string {
	key := []byte(accessKeySecret + "&")
	h := hmac.New(sha1.New, key)
	h.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
