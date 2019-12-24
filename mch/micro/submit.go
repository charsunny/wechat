package micro

import (
	"encoding/xml"
	"github.com/charsunny/wechat/mch/core"
	wechatutil "github.com/charsunny/wechat/util"
)

// Submit 申请入驻接口提交你的小微商户资料，申请后一般5分钟左右可以查询到具体的申请结果
//  NOTE: 请求需要双向证书.
func Submit(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/applyment/micro/submit", req)
}

type SubmitRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数, TransactionId 和 OutTradeNo 二选一即可.
	CertSn string `xml:"cert_sn"` 				// 平台证书序列号
	BusinessCode    string `xml:"business_code"`   // 服务商自定义的商户唯一编号。每个编号对应一个申请单，每个申请单审核通过后会生成一个微信支付商户号。

	IdCardCopy   string `xml:"id_card_copy"`  // 身份证人像面照片，请填写由图片上传接口预先上传图片生成好的media_id
	IdCardNational      string  `xml:"id_card_national"`      // 身份证国徽面照片，请填写由图片上传接口预先上传图片生成好的media_id
	IdCardName     string  `xml:"id_card_name"`     // 身份证姓名, 需加密处理
	IdCardNumber     string  `xml:"id_card_number"`     // 身份证号码，需加密处理
	IdCardValidTime    string  `xml:"id_card_valid_time"`     // 身份证有效期限

	AccountName     string  `xml:"account_name"`     // 必须与身份证姓名一致，加密处理
	AccountBank     string  `xml:"account_bank"`     // 开户银行
	BankAddressCode    string  `xml:"bank_address_code"`     // 开户银行省市编码
	BankName     string  `xml:"bank_name"`     // 开户银行全称（含支行）1）17家直连银行无需填写，其他银行请务必填写 2）需填写银行全称，如"深圳农村商业银行XXX支行"
	AccountNumber    string  `xml:"account_number"`     // 银行账号 数字，长度遵循系统支持的对私卡号长度要求,该字段需进行加密处理，小微商户开户目前不支持以下前缀的银行卡"623501;621468;620522;625191;622384;623078;940034;622150;622151;622181;622188;955100;621095;620062;621285;621798;621799;621797;622199;621096;62215049;62215050;62215051;62218849;62218850;62218851;621622;623219;621674;623218;621599;623698;623699;623686;621098;620529;622180;622182;622187;622189;621582;623676;623677;622812;622810;622811;628310;625919;625368;625367;518905;622835;625603;625605;518905"

	StoreName     string  `xml:"store_name"`     // 门店名称，最长50个中文字符
	StoreAddressCode     string  `xml:"store_address_code"`     // 门店场所：填写门店省市编码 流动经营/便民服务：填写经营/服务所在地省市编码 线上商品/服务交易：填写卖家所在地省市编码
	StoreStreet    string  `xml:"store_street"`     // 门店街道名称 门店场所：填写店铺详细地址，具体区/县及街道门牌号或大厦楼层 流动经营/便民服务：填写“无" 线上商品/服务交易：填写电商平台名称
	StoreEntrancePic   string `xml:"store_entrance_pic"`  // 门店场所：提交门店门口照片，要求招牌清晰可见 流动经营/便民服务：提交经营/服务现场照片 线上商品/服务交易：提交店铺首页截图
	IndoorPic      string  `xml:"indoor_pic"`      // 门店场所：提交店内环境照片 流动经营/便民服务：可提交另一张经营/服务现场照片 线上商品/服务交易：提交店铺管理后台截图
	MerchantShortname     string  `xml:"merchant_shortname"`     // 最多16个汉字长度， 需与商家的实际经营场景相符
	ServicePhone     string  `xml:"service_phone"`			// 客服电话
	ProductDesc     string  `xml:"product_desc"`			// 请填写以下描述之一：餐饮 线下零售 居民生活服务 休闲娱乐 交通出行 其他
	Rate  string  `xml:"rate"`								// 费率 由服务商指定，微信支付提供字典值 枚举值0.38%、0.39%、0.4%、0.45%、0.48%、0.49%、0.5%、0.55%、0.58%、0.59%、0.6%
	Contact     string  `xml:"contact"`			// 超级管理员姓名 请确定其为商户法定代表人或负责人, 该字段需进行加密处理，加密方法详见敏感信息加密说明
	ContactPhone     string  `xml:"contact_phone"`			// 11位数字，手机号码 ，该字段需进行加密处理，加密方法详见敏感信息加密说明

	// 可选参数
	StoreLongitude     string  `xml:"store_longitude"`     // 门店经度
	StoreLatitude    string  `xml:"store_latitude"`     // 门店纬度
	AddressCertification      string  `xml:"address_certification"`   // 门面租赁合同扫描件或经营场地证明（需与身份证同名）
	BusinessAdditionDesc      string  `xml:"business_addition_desc"`   // 可填写需要额外说明的文字
	BusinessAdditionPics      string  `xml:"business_addition_pics"`   // 最多可上传5张照片，请填写已预先上传图片生成好的MediaID
	ContactEmail     string  `xml:"contact_email"`			// 需要带@，遵循邮箱格式校验 ，该字段需进行加密处理，加密方法详见敏感信息加密说明
	NonceStr      string `xml:"nonce_str"`       // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType      string `xml:"sign_type"`       // 签名类型，仅支持HMAC-SHA256 。该字段需参与签名sign的计算。
	Version	 	string `xml:"version"`			// 固定版本号为3.0
}

type SubmitResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	ApplymentId string `xml:"applyment_id"` // 微信订单号
}

// Submit2 申请入驻接口提交你的小微商户资料，申请后一般5分钟左右可以查询到具体的申请结果
//  NOTE:
//  1. 请求需要双向证书.

func Submit2(clt *core.Client, req *SubmitRequest) (resp *SubmitResponse, err error) {
	m1 := make(map[string]string, 16)

	if req.NonceStr == "" {
		req.NonceStr = wechatutil.NonceStr()
	}
	req.SignType = core.SignType_HMAC_SHA256
	req.Version = "3.0"

	data, err :=  xml.Marshal(req)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, &m1)
	if err != nil {
		return
	}
	m2, err := Submit(clt, m1)

	if err != nil {
		return nil, err
	}

	resp = &SubmitResponse{
		ApplymentId: m2["applyment_id"],
	}

	return resp, nil
}
