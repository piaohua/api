package jtpay

import (
	"errors"
	"strings"
)

// 下单请求参数
type JTpayOrder struct {
	P1_usercode           string `json:"p1_usercode"`
	P2_order              string `json:"p2_order"`
	P3_money              string `json:"p3_money"`
	P4_returnurl          string `json:"p4_returnurl"`
	P5_notifyurl          string `json:"p5_notifyurl"`
	P6_ordertime          string `json:"p6_ordertime"`
	P7_sign               string `json:"p7_sign"`
	P8_signtype           string `json:"p8_signtype"`
	P9_paymethod          string `json:"p9_paymethod"`
	P10_paychannelnum     string `json:"p10_paychannelnum"`
	P11_cardtype          string `json:"p11_cardtype"`
	P12_channel           string `json:"p12_channel"`
	P13_orderfailertime   string `json:"p13_orderfailertime"`
	P14_customname        string `json:"p14_customname"`
	P15_customcontacttype string `json:"p15_customcontacttype"`
	P16_customcontact     string `json:"p16_customcontact"`
	P17_customip          string `json:"p17_customip"`
	P18_product           string `json:"p18_product"`
	P19_productcat        string `json:"p19_productcat"`
	P20_productnum        string `json:"p20_productnum"`
	P21_pdesc             string `json:"p21_pdesc"`
	P22_version           string `json:"p22_version"`
	P23_charset           string `json:"p23_charset"`
	P24_remark            string `json:"p24_remark"`
	P25_terminal          string `json:"p25_terminal"`
	P26_iswappay          string `json:"p26_iswappay"`
}

/*
p1_usercode                 string() 竣付通分配的商户号。 必填
p2_order                 string(29) 用户订单号,建议 8 位商户号+14 位时间 yyyymmddhhmmss+5 位 流水号,中间用“-”分隔。例如: 12345678-20150728132430-12 345(只是建议,商户的订单号也 可以不采用这种格式)。 必填
p3_money                 string(18) 订单金额,精确到分。例如 99.99。 必填
p4_returnurl                 string(190) 用户明文跳转地址,用于告知付款 人支付结果。必须包含 http://或 https://。注意:该 URL 建议不包 含 GET 参数,即形如“? name=value”的内容,竣付通支 付网关不确保这些 GET 参数通过后 台方式向商户反馈时能被保留。 必填
p5_notifyurl                 string(190) 服务器异步通知地址,用于通知商 户系统处理业务(数据库更新等)。 必须包含 http://或 https://。注意: 该 URL 建议不包含 GET 参数,即 形如“?name=value”的内容, 竣付通支付网关不确保这些 GET 参 数通过后台方式向商户反馈时能被 保留。 必填
p6_ordertime                 string(14) 商户订单时间,格式 yyyymmddhhmmss。 必填
p7_sign                 string(256) 商户传递参数加密值,约定 p1_ usercode + "&" +p2_ order + "&" +p3_ money + "&" +p4_returnurl + "&" +p5_ noticeurl + "&" +p6_ordertime +CompKey 连接起来进行 MD5 加密后 32 位大写字符串,(参数之 间必须添加&符号,最后 p6_ordertime 和 CompKey 之间 不加&符号。CompKey 为商户的 秘钥)目前只限定 md5 加密。 必填
p8_signtype                 string(5) 签名验证方式:1、MD5,传固定 值 1。如果用户传递参数为空,则 默认 MD5 验证。 可空
p9_paymethod                 string(5) 商户支付方式:固定值 3。如果用 户传递参数为空,则默认网银支付。 可空
p10_paychannelnum                 string(12) 支付通道编码。 可空
p11_cardtype                 string(5) 为空 可空
p12_channel                 string(5) 为空 可空
p13_orderfailertime                 string(14) 订单失效时间,格式为 14 位时间格 式:yyyymmddhhmmss,精确到秒,超时则此订单失效。 可空
p14_customname                 string(128) 客户、或者玩家所在平台账号。请 务必填写真实信息,否则将影响后 续查单结果。 必填
p15_customcontacttype                 string(10) 客户联系方式:1、email,2、 phone,3、地址。 可空
p16_customcontact                 string(200) 客户联系方式。 可空
p17_customip                 string(128) 客户 ip 地址,规定以 192_168_0_253 格式,如果以 “192.168.0.253”可能会发生签名 错误。 必填
p18_product                 string(256) 商品名称。 可空
p19_productcat                 string(200) 商品种类。 可空
p20_productnum                 string(10) 商品数量,不传递参数默认 0。 可空
p21_pdesc                 string(200) 商品描述。 可空
p22_version                 string(5) 接口版本,目前默认 2.0。 可空
p23_charset                 string(5) 提交的编码格式,1、UTF-8,2、 GBK/GB2312,默认 UTF-8。 可空
p24_remark                 string(256) 备注。此参数我们会在下行过程中 原样返回。您可以在此参数中记录 一些数据,方便在下行过程中直接 读取。 可空
p25_terminal                 string(5) 终端设备固定值 2。 必填
p26_iswappay                 string(5) 支付场景固定值 3。 必填
*/

// 交易结果通知
type NotifyResult struct {
	P1_usercode      string `json:"p1_usercode"`
	P2_order         string `json:"p2_order"`
	P3_money         string `json:"p3_money"`
	P4_status        string `json:"p4_status"`
	P5_jtpayorder    string `json:"p5_jtpayorder"`
	P6_paymethod     string `json:"p6_paymethod"`
	P7_paychannelnum string `json:"p7_paychannelnum"`
	P8_charset       string `json:"p8_charset"`
	P9_signtype      string `json:"p9_signtype"`
	P10_sign         string `json:"p10_sign"`
	P11_remark       string `json:"p11_remark"`
}

//// Parse the reponse message
//func ParseNotifyResult(resp []byte) (*NotifyResult, error) {
//	result := new(NotifyResult)
//	err := json.Unmarshal(resp, result)
//	if err != nil {
//		return result, err
//	}
//	return result, nil
//}

// ParseNotifyResult convert the string to struct
func ParseNotifyResult(resp []byte) (*NotifyResult, error) {
	result := new(NotifyResult)
	s := strings.Split(string(resp), "&")
	for _, v := range s {
		args := strings.Split(v, "=")
		if len(args) != 2 {
			return nil, errors.New("args error")
		}
		switch args[0] {
		case "p1_usercode":
			result.P1_usercode = args[1]
		case "p2_order":
			result.P2_order = args[1]
		case "p3_money":
			result.P3_money = args[1]
		case "p4_status":
			result.P4_status = args[1]
		case "p5_jtpayorder":
			result.P5_jtpayorder = args[1]
		case "p6_paymethod":
			result.P6_paymethod = args[1]
		case "p7_paychannelnum":
			result.P7_paychannelnum = args[1]
		case "p8_charset":
			result.P8_charset = args[1]
		case "p9_signtype":
			result.P9_signtype = args[1]
		case "p10_sign":
			result.P10_sign = args[1]
		case "p11_remark":
			result.P11_remark = args[1]
		default:
			//TODO
		}
	}
	return result, nil
}

/*
p1_usercode 用户编号 文本 (必须)商户在竣付通的商户号。
p2_order 订单号 文本 (必须)商户提交的订单号。
p3_money 金额 货币 (必须)该次交易金额(以该金额为准)。
p4_status 支付结果 文本 (必须)支付返回结果 1 代表成功,其 他为失败。
p5_jtpayorder 竣付通订单号 文本 (必须)返回竣付通的订单号。
p6_paymethod 商户支付方式 文本 (必须)订单的支付方式。
p7_paychannelnum 支付通道编码 文本 微信不回调此参数。
p8_charset 编码格式 文本 (必须)商户提交订单时候传递的编码 格式。
p9_signtype 签名验证方式 文本 (必须)签名验证方式。
p10_sign 验证签名 文本 MD5(p1_usercode&p2_order&p3_ money&p4_status&p5_jtpayorder &p6_paymethod&&p8_charset&p 9_signtype&CompKey)。各字段中间 添加“&”分隔符,区分不同字段进行 加密验签。注意最后“p9_signtype &MD5key“中间有“&”符号,加密 形式为 MD5 32 位大写字符串。
p11_remark string(256) 文本 备注,原样返回用户提交的 p24_remark 信息
*/
