package jtpay

import (
	"fmt"
	"testing"
)

func TestPay(t *testing.T) {
	cfg := &JTpayConfig{
		Appid:         "01018035867501",
		CompKey:       "040109141916ze44Mfoy",
		PlaceOrderUrl: "http://order.z.jtpay.com/jh-web-order/order/receiveOrder",
	}
	trans, err := NewJTpayTrans(cfg)
	if err != nil {
		t.Logf("err %v", err)
	}
	order := &JTpayOrder{
		P3_money:       "1",
		P14_customname: "101409",
		P17_product:    "3",
		P19_productnum: "1",
		P25_terminal:   "2",
		P7_productcode: "WX",
		P16_customip:   "192_168_0_253",
	}
	order.P1_yingyongnum = trans.Config.Appid
	t.Logf("appid %s", order.P1_yingyongnum)
	//order.P6_ordertime = OrderTimeStr()
	//order.P2_ordernumber = GenOrderid(order.P1_yingyongnum, order.P6_ordertime)
	//E5080E76A63C6803A3D8A4F8D1AE8CA2
	//E5080E76A63C6803A3D8A4F8D1AE8CA2
	trans.InitOrder(order)
	resp, err := trans.Submit(order)
	t.Logf("resp %s, err %v", resp, err)
	t.Logf("err %v", err)
	fmt.Printf("%s\n", order.P8_sign)
	t.Logf("order %#v", order)
}

func TestVerify(t *testing.T) {
	cfg := &JTpayConfig{
		Appid:         "",
		CompKey:       "",
		PlaceOrderUrl: "http://order.z.jtpay.com/jh-web-order/order/receiveOrder",
	}
	trans, err := NewJTpayTrans(cfg)
	if err != nil {
		panic(err)
	}
	//result := &NotifyResult{
	//	P1_yingyongnum:    "10231993",
	//	P2_ordernumber:    "10231993-20180327000845-48314",
	//	P3_money:          "1",
	//	P4_zfstate:        "1",
	//	P5_orderid:        "20180326160844234406",
	//	P6_productcode:    "3",
	//	P7_bank_card_code: "",
	//	P8_charset:        "UTF-8",
	//	P9_signtype:       "MD5",
	//	P10_sign:          "EC03C50E252D2A3CC550130243A540E9",
	//	P11_pdesc:         "",
	//}
	//
	//t.Log(trans.NotifyVerify(result))
	//
	notify := &NotifyResult{
		P1_yingyongnum:    "01018076485101",
		P2_ordernumber:    "01018076485101-20180713150117-35583",
		P3_money:          "12.00",
		P4_zfstate:        "1",
		P5_orderid:        "0102018071315011700916",
		P6_productcode:    "ZFBZZWAP",
		P7_bank_card_code: "",
		P8_charset:        "UTF-8",
		P9_signtype:       "1",
		P10_sign:          "DECEC550FC4ABB915CB98EAC66D8334A",
		P11_pdesc:         "",
		P12_remark:        "",
		P13_zfmoney:       "12.00"}
	t.Log(trans.NotifyVerify(notify))
}

//下单页面
func ioswapHtml(order *JTpayOrder) (str string) {
	str = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>竣付通</title>
</head>
<!--支付宝IOSwap支付请求提交页-->
<body onLoad="document.yeepay.submit();">
	<form name='yeepay' action='http://order.z.jtpay.com/jh-web-order/order/receiveOrder' method='post'  >
	`
	str += fmt.Sprintf("<input type='hidden' name='p1_yingyongnum'				value='%s'>", order.P1_yingyongnum)
	str += fmt.Sprintf("<input type='hidden' name='p2_ordernumber'				value='%s'>      ", order.P2_ordernumber)
	str += fmt.Sprintf("<input type='hidden' name='p3_money'				value='%s'>      ", order.P3_money)
	str += fmt.Sprintf("<input type='hidden' name='p6_ordertime'			value='%s'>  ", order.P6_ordertime)
	str += fmt.Sprintf("<input type='hidden' name='p7_productcode'					value='%s'>       ", order.P7_productcode)
	str += fmt.Sprintf("<input type='hidden' name='p8_sign'					value='%s'>       ", order.P8_sign)
	str += fmt.Sprintf("<input type='hidden' name='p9_signtype'			value='%s'>  ", order.P9_signtype)
	str += fmt.Sprintf("<input type='hidden' name='p25_terminal'			value='%s'>  ", order.P25_terminal)
	str += `
	</form>
</body>
</html>
`
	return
}
