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

//GenOrderid 订单号
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

//p1_yingyongnum + "&" +p2_ordernumber + "&" +p3_money + "&" +p6_ordertime + "&"
//+p7_productcode + "&" +CompKey 连接起来进行 MD5 加密后 32 位大写字符串,
func p8sign(o *JTpayOrder, CompKey string) (s string) {
	s = o.P1_yingyongnum + "&" + o.P2_ordernumber + "&" + o.P3_money + "&" +
		o.P6_ordertime + "&" + o.P7_productcode + "&" + CompKey
	return Md5(s)
}

/*
$p10_sign=strtoupper(md5($p1_yingyongnum."&".
$p2_ordernumber."&".
$p3_money."&".
$p4_zfstate."&".
$p5_orderid."&".
$p6_productcode."&".
$p7_bank_card_code."&".
$p8_charset."&".
$p9_signtype."&".
$p11_pdesc."&".
$p13_zfmoney."&".
$compkey));
*/
func p10sign(r *NotifyResult, CompKey string) (s string) {
	s = r.P1_yingyongnum + "&" + r.P2_ordernumber + "&" + r.P3_money + "&" +
		r.P4_zfstate + "&" + r.P5_orderid + "&" + r.P6_productcode + "&" +
		r.P7_bank_card_code + "&" + r.P8_charset + "&" + r.P9_signtype + "&" +
		r.P11_pdesc + "&" + r.P13_zfmoney + "&" + CompKey
	return ToUpper(Md5(s))
}

//ip地址
func customip(ip string) string {
	s := strings.Split(ip, ".")
	return strings.Join(s, "_")
}

//OrderTimeStr 下单时间
func OrderTimeStr() string {
	return time.Now().Format("20060102150405")
}

//ToUpper golang make the caracter in a string uppercase
func ToUpper(str string) string {
	var s string
	for _, v := range str {
		s += string(unicode.ToUpper(v))
	}
	return s
}

// Md5 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}