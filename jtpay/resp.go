package jtpay

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//订单号
func GenOrderid(usercode, ordertime string) (s string) {
	t := time.Now().UnixNano()
	r := rand.NewSource(t)
	o := rand.New(r)
	s += usercode + "-"
	s += ordertime + "-"
	var i int = 5
	for i > 0 {
		s += strconv.FormatInt(o.Int63n(10), 10)
		i--
	}
	return
}

//p1_ usercode + "&" +p2_ order + "&" +p3_ money + "&" +p4_returnurl + "&"
//+p5_ noticeurl + "&" +p6_ordertime +CompKey 连接起来进行 MD5 加密后 32 位大写字符串,
//(参数之 间必须添加&符号,最后 p6_ordertime 和 CompKey 之间 不加&符号
//。CompKey 为商户的 秘钥)目前只限定 md5 加密。 必填
func p7sign(p1_usercode, p2_order, p3_money, p4_returnurl,
	p5_noticeurl, p6_ordertime, CompKey string) (s string) {
	s = p1_usercode + "&" + p2_order + "&" + p3_money + "&" +
		p4_returnurl + "&" + p5_noticeurl + "&" + p6_ordertime + CompKey
	return ToUpper(Md5(s))
}

//p10_sign 验证签名 文本 MD5(p1_usercode&p2_order&p3_ money&p4_status&p5_jtpayorder &p6_paymethod&&p8_charset&p 9_signtype&CompKey)。
//各字段中间 添加“&”分隔符,区分不同字段进行 加密验签。注意最后“p9_signtype &MD5key“中间有“&”符号,加密 形式为 MD5 32 位大写字符串。
func p10sign(p1_usercode, p2_order, p3_money, p4_status, p5_jtpayorder,
	p6_paymethod, p8_charset, p9_signtype, CompKey string) (s string) {
	s = p1_usercode + "&" + p2_order + "&" + p3_money + "&" +
		p4_status + "&" + p5_jtpayorder + "&" + p6_paymethod + "&&" +
		p8_charset + "&" + p9_signtype + "&" + CompKey
	return ToUpper(Md5(s))
}

//ip地址
func customip(ip string) string {
	s := strings.Split(ip, ".")
	return strings.Join(s, "_")
}

//下单时间
func OrderTimeStr() string {
	return time.Now().Format("20060102150405")
}

//golang make the caracter in a string uppercase
func ToUpper(str string) string {
	var s string
	for _, v := range str {
		s += string(unicode.ToUpper(v))
	}
	return s
}

// md5 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
