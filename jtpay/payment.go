package jtpay

import (
	"encoding/json"
	"errors"
	"strings"
)

//JTpayOrder 下单请求参数
type JTpayOrder struct {
	P1_yingyongnum      string `json:"p1_yingyongnum"`
	P2_ordernumber      string `json:"p2_ordernumber"`
	P3_money            string `json:"p3_money"`
	P6_ordertime        string `json:"p6_ordertime"`
	P7_productcode      string `json:"p7_productcode"`
	P8_sign             string `json:"p8_sign"`
	P9_signtype         string `json:"p9_signtype"`
	P10_bank_card_code  string `json:"p10_bank_card_code"`
	P11_cardtype        string `json:"p11_cardtype"`
	P12_channel         string `json:"p12_channel"`
	P13_orderfailertime string `json:"p13_orderfailertime"`
	P14_customname      string `json:"p14_customname"`
	P15_customcontact   string `json:"p15_customcontact"`
	P16_customip        string `json:"p16_customip"`
	P17_product         string `json:"p17_product"`
	P18_productcat      string `json:"p18_productcat"`
	P19_productnum      string `json:"p19_productnum"`
	P20_pdesc           string `json:"p20_pdesc"`
	P21_version         string `json:"p21_version"`
	P22_sdkversion      string `json:"p22_sdkversion"`
	P23_charset         string `json:"p23_charset"`
	P24_remark          string `json:"p24_remark"`
	P25_terminal        string `json:"p25_terminal"`
	P26_ext1            string `json:"p26_ext1"`
	P27_ext2            string `json:"p27_ext2"`
	P28_ext3            string `json:"p28_ext3"`
	P29_ext4            string `json:"p29_ext4"`
}

/*
$compkey 		    = "040109141916ze44Mfoy";	//商户密钥
$p1_yingyongnum	    = "01018035867501";			//商户应用号
$p2_ordernumber     = "PHP".date("YmdHis", time());	//商户订单号
$p3_money 		    = $_POST['p3_money'];		//商户订单金额，保留两位小数
$p6_ordertime  	    = date("YmdHis", time());	//商户订单时间
$p7_productcode	    = "ZFBZFWAP";				//产品支付类型编码
$presign 		    = $p1_yingyongnum."&".$p2_ordernumber."&".$p3_money."&".$p6_ordertime."&".$p7_productcode."&".$compkey;
$p8_sign 		    = md5($presign);		    //订单签名
$p9_signtype 	    = "1";					    //签名方式
$p10_bank_card_code = "";						//银行卡或卡类编码
$p11_cardtype  	    = "";						//商户支付银行卡类型id
$p12_channel 	    = "";						//商户支付银行卡类型长度
$p13_orderfailertime= "";						//订单失效时间
$p14_customname     = $_POST['p14_customname']; //商户游戏账号
$p15_customcontact  = "";						//商户联系内容
$p16_customip  	    = "192_168_0_253";			//付款ip地址
$p17_product  	    = "xxx";					//商户名称
$p18_productcat	    = "";						//商品种类
$p19_productnum     = "";						//商品数量
$p20_pdesc          = "";						//商品描述
$p21_version        = "";						//对接版本
$p22_sdkversion	    = "";						//SDK版本
$p23_charset   	    = "UTF-8";					//编码格式
$p24_remark    	    = "";						//备注
$p25_terminal  	    = "2";				        //商户终端设备值
// 终端设备值1 pc 2 ios  3 安卓
$p26_ext1     		= ""; 				        //预留参数
$p27_ext2    		= "";				        //预留参数
$p28_ext3     		= "";				        //预留参数
$p29_ext4     		= "";				        //预留参数
*/

//NotifyResult 交易结果通知
type NotifyResult struct {
	P1_yingyongnum    string `json:"p1_yingyongnum"`
	P2_ordernumber    string `json:"p2_ordernumber"`
	P3_money          string `json:"p3_money"`
	P4_zfstate        string `json:"p4_zfstate"`
	P5_orderid        string `json:"p5_orderid"`
	P6_productcode    string `json:"p6_productcode"`
	P7_bank_card_code string `json:"p7_bank_card_code"`
	P8_charset        string `json:"p8_charset"`
	P9_signtype       string `json:"p9_signtype"`
	P10_sign          string `json:"p10_sign"`
	P11_pdesc         string `json:"p11_pdesc"`
	P12_remark        string `json:"p12_remark"`
	P13_zfmoney       string `json:"p13_zfmoney"`
}

// ParseNotifyResult2 Parse the reponse message
func ParseNotifyResult2(resp []byte) (*NotifyResult, error) {
	result := new(NotifyResult)
	err := json.Unmarshal(resp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

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
		case "p1_yingyongnum":
			result.P1_yingyongnum = args[1]
		case "p2_ordernumber":
			result.P2_ordernumber = args[1]
		case "p3_money":
			result.P3_money = args[1]
		case "p4_zfstate":
			result.P4_zfstate = args[1]
		case "p5_orderid":
			result.P5_orderid = args[1]
		case "p6_productcode":
			result.P6_productcode = args[1]
		case "p7_bank_card_code":
			result.P7_bank_card_code = args[1]
		case "p8_charset":
			result.P8_charset = args[1]
		case "p9_signtype":
			result.P9_signtype = args[1]
		case "p10_sign":
			result.P10_sign = args[1]
		case "p11_pdesc":
			result.P11_pdesc = args[1]
		case "p12_remark":
			result.P12_remark = args[1]
		case "p13_zfmoney":
			result.P13_zfmoney = args[1]
		default:
			//TODO
		}
	}
	return result, nil
}

/*
p1_yingyongnum 用户编号 Y 商户在竣付通平台的应用 ID。
p2_ordernumber 订单号 Y 商户提交的订单号。
p3_money 金额 Y 该次交易金额（以通知金额为准）。无论 商户发起支付时金额采用哪种格式，返回 金额均保留两位小数。
p4_zfstate 支付结果 Y （必须）支付返回结果 1 代表成功，其他 为失败。
p5_orderid 竣付通订单号 Y （必须）返回竣付通的订单号。
p6_productcode 商户支付方式 Y （必须）订单的支付方式。
p7_bank_card_code 支付通道编码 N
p8_charset 编码格式 Y （必须）商户提交订单时候传递的编码格 式
p9_signtype 签名验证方式 Y （必须）签名验证方式。
p10_sign 验证签名 Y 格式为 32 位大写字符串，签名算法见 5.2 签名算法
p11_pdesc string（256） Y 备注，原样返回用户提交的 p20_pdesc 信息
p12_remark string（256） Y 备注，原样返回用户提交的 p24_remark 信息
p13_zfmoney 金额 Y 实际支付金额，保留两位小数
*/
