package applyment4sub

import (
	"encoding/json"

	core "github.com/charsunny/wechat/mch/core_v3"
)

type ApplymentReq struct {
	BusinessCode string `json:"business_code"` // 业务申请编号
	ContactInfo  struct {
		ContactName     string `json:"contact_name"`                // 超级管理员姓名
		ContactIdNumber string `json:"contact_id_number,omitempty"` // 超级管理员身份证件号码
		Openid          string `json:"openid,omitempty"`            // 超级管理员微信openid
		MobilePhone     string `json:"mobile_phone"`                // 联系手机
		ContactEmail    string `json:"contact_email,omitempty"`     // 邮箱
	} `json:"contact_info"` // 超管
	SubjectInfo struct {
		SubjectType         string `json:"subject_type"` // 主体类型
		BusinessLicenseInfo struct {
			LicenseCopy   string `json:"license_copy"`   // 营业执照照片
			LicenseNumber string `json:"license_number"` // 注册号/统一社会信用代码
			MerchantName  string `json:"merchant_name"`  // 商户名称
			LegalPerson   string `json:"legal_person"`   // 个体户经营者/法人姓名
		} `json:"business_license_info,omitempty"` // 营业执照
		CertificateInfo struct {
			CertCopy       string `json:"cert_copy"`       // 登记证书照片
			CertType       string `json:"cert_type"`       // 登记证书类型
			CertNumber     string `json:"cert_number"`     // 证书号
			MerchantName   string `json:"merchant_name"`   // 商户名称
			CompanyAddress string `json:"company_address"` // 注册地址
			LegalPerson    string `json:"legal_person"`    // 法定代表人
			PeriodBegin    string `json:"period_begin"`    // 有效期限开始日期
			PeriodEnd      string `json:"period_end"`      // 有效期限结束日期
		} `json:"certificate_info,omitempty"` // 登记证书
		OrganizationInfo struct {
			OrganizationCopy string `json:"organization_copy"` // 组织机构代码证照片
			OrganizationCode string `json:"organization_code"` // 组织机构代码
			OrgPeriodBegin   string `json:"org_period_begin"`  // 组织机构代码证有效期开始日期
			OrgPeriodEnd     string `json:"org_period_end"`    // 组织机构代码证有效期结束日期
		} `json:"organization_info,omitempty"` // 组织机构代码证
		CertificateLetterCopy string `json:"certificate_letter_copy,omitempty"` // 单位证明函照片
		IdentityInfo          struct {
			IdDocType  string `json:"id_doc_type"` // 证件类型
			IdCardInfo struct {
				IdCardCopy      string `json:"id_card_copy"`      // 身份证人像面照片
				IdCardNational  string `json:"id_card_national"`  // 身份证国徽面照片
				IdCardName      string `json:"id_card_name"`      // 身份证姓名
				IdCardNumber    string `json:"id_card_number"`    // 身份证号码
				CardPeriodBegin string `json:"card_period_begin"` // 身份证有效期开始时间
				CardPeriodEnd   string `json:"card_period_end"`   // 身份证有效期结束时间
			} `json:"id_card_info,omitempty"` // 身份证信息
			IdDocInfo struct {
				IdDocCopy      string `json:"id_doc_copy"`      // 证件照片
				IdDocName      string `json:"id_doc_name"`      // 证件姓名
				IdDocNumber    string `json:"id_doc_number"`    // 证件号码
				DocPeriodBegin string `json:"doc_period_begin"` // 证件有效期开始时间
				DocPeriodEnd   string `json:"doc_period_end"`   // 证件有效期结束时间
			} `json:"id_doc_info,omitempty"` // 其他类型证件信息
			Owner bool `json:"owner"` // 经营者/法人是否为受益人
		} `json:"identity_info"` // 经营者/法人身份证件
		UboInfo struct {
			IdType         string `json:"id_type"`                    // 证件类型
			IdCardCopy     string `json:"id_card_copy,omitempty"`     // 身份证人像面照片
			IdCardNational string `json:"id_card_national,omitempty"` // 身份证国徽面照片
			IdDocCopy      string `json:"id_doc_copy,omitempty"`      // 证件照片
			Name           string `json:"name"`                       // 受益人姓名
			IdNumber       string `json:"id_number"`                  // 证件号码
			IdPeriodBegin  string `json:"id_period_begin"`            // 证件有效期开始时间
			IdPeriodEnd    string `json:"id_period_end"`              // 证件有效期结束时间
		} `json:"ubo_info,omitempty"` // 最终受益人信息
	} `json:"subject_info"` // 主体资料
	BusinessInfo struct {
		MerchantShortname string `json:"merchant_shortname"` // 商户简称
		ServicePhone      string `json:"service_phone"`      // 客服电话
		SalesInfo         struct {
			SalesScenesType string `json:"sales_scenes_type"` // 经营场景类型
			BizStoreInfo    struct {
				BizStoreName     string `json:"biz_store_name"`     // 门店名称
				BizAddressCode   string `json:"biz_address_code"`   // 门店省市编码
				BizStoreAddress  string `json:"biz_store_address"`  // 门店地址
				StoreEntrancePic string `json:"store_entrance_pic"` // 门店门头照片
				IndoorPic        string `json:"indoor_pic"`         // 店内环境照片
				BizSubAppid      string `json:"biz_sub_appid"`      // 线下场所对应的商家APPID
			} `json:"biz_store_info,omitempty"` // 线下门店场景
			MpInfo struct {
				MpAppid    string `json:"mp_appid,omitempty"`     // 服务商公众号APPID
				MpSubAppid string `json:"mp_sub_appid,omitempty"` // 商家公众号APPID
				MpPics     string `json:"mp_pics"`                // 公众号页面截图
			} `json:"mp_info,omitempty"` // 公众号场景
			MiniProgramInfo struct {
				MiniProgramAppid    string `json:"mini_program_appid,omitempty"`     // 服务商小程序APPID
				MiniProgramSubAppid string `json:"mini_program_sub_appid,omitempty"` // 商家小程序APPID
				MiniProgramPics     string `json:"mini_program_pics"`                // 小程序截图
			} `json:"mini_program_info,omitempty"` // 小程序场景
			AppInfo struct {
				AppAppid    string `json:"app_appid,omitempty"`     // 服务商应用APPID
				AppSubAppid string `json:"app_sub_appid,omitempty"` // 商家应用APPID
				AppPics     string `json:"app_pics"`                // APP截图
			} `json:"app_info,omitempty"` // APP场景
			WebInfo struct {
				Domain           string `json:"domain"`                      // 互联网网站域名
				WebAuthorisation string `json:"web_authorisation,omitempty"` // 网站授权函
				WebAppid         string `json:"web_appid,omitempty"`         // 互联网网站对应的商家APPID
			} `json:"web_info,omitempty"` // 互联网网站场景
			WeworkInfo struct {
				SubCorpId  string `json:"sub_corp_id"` // 商家企业微信CorpID
				WeworkPics string `json:"wework_pics"` // 企业微信页面截图
			} `json:"wework_info,omitempty"` // 企业微信场景
		} `json:"sales_info"` // 经营场景
	} `json:"business_info"` // 经营资料
	SettlementInfo struct {
		SettlementId        string `json:"settlement_id"`                  // 入驻结算规则ID
		QualificationType   string `json:"qualification_type"`             // 所属行业
		Qualifications      string `json:"qualifications,omitempty"`       // 特殊资质图片
		ActivitiesId        string `json:"activities_id,omitempty"`        // 优惠费率活动ID
		ActivitiesRate      string `json:"activities_rate,omitempty"`      // 优惠费率活动值
		ActivitiesAdditions string `json:"activities_additions,omitempty"` // 优惠费率活动补充材料
	} `json:"settlement_info"` // 结算规则
	BankAccountInfo struct {
		BankAccountType string `json:"bank_account_type"`        // 账户类型
		AccountName     string `json:"account_name"`             // 开户名称
		AccountBank     string `json:"account_bank"`             // 开户银行
		BankAddressCode string `json:"bank_address_code"`        // 开户银行省市编码
		BankBranchId    string `json:"bank_branch_id,omitempty"` // 开户银行联行号
		BankName        string `json:"bank_name,omitempty"`      // 开户银行全称（含支行)
		AccountNumber   string `json:"account_number"`           // 银行账号
	} `json:"bank_account_info"` // 结算银行账户
	AdditionInfo struct {
		LegalPersonCommitment string `json:"legal_person_commitment,omitempty"` // 法人开户承诺函
		LegalPersonVideo      string `json:"legal_person_video,omitempty"`      // 法人开户意愿视频
		BusinessAdditionPics  string `json:"business_addition_pics,omitempty"`  // 补充材料
		BusinessAdditionMsg   string `json:"business_addition_msg,omitempty"`   // 补充说明
	} `json:"addition_info,omitempty"` // 补充材料
}

type ApplymentReply struct {
	ApplymentId string `json:"applyment_id"`
}

type ApplymentQueryReply struct {
	BusinessCode      string `json:"business_code"`       // 业务申请编号
	ApplymentId       string `json:"applyment_id"`        // 微信支付申请单号
	SubMchid          string `json:"sub_mchid"`           // 特约商户号
	SignUrl           string `json:"sign_url"`            // 超级管理员签约链接
	ApplymentState    string `json:"applyment_state"`     // 申请单状态
	ApplymentStateMsg string `json:"applyment_state_msg"` // 申请状态描述
	AuditDetail       struct {
		Field        string `json:"field"`         // 字段名
		FieldName    string `json:"field_name"`    // 字段名称
		RejectReason string `json:"reject_reason"` // 驳回原因
	} `json:"audit_detail"` // 驳回原因详情
}

type ModifySettlementReq struct {
	SubMchid        string `json:"sub_mchid"`                // 特约商户号
	AccountType     string `json:"account_type"`             // 账户类型
	AccountBank     string `json:"account_bank"`             // 开户银行
	BankAddressCode string `json:"bank_address_code"`        // 开户银行省市编码
	BankName        string `json:"bank_name,omitempty"`      // 开户银行全称（含支行）
	BankBranchId    string `json:"bank_branch_id,omitempty"` // 开户银行联行号
	AccountNumber   string `json:"account_number"`           // 银行账号
}

type QuerySettlementReply struct {
	AccountType   string `json:"account_type"`   // 账户类型
	AccountBank   string `json:"account_bank"`   // 开户银行
	BankName      string `json:"bank_name"`      // 开户银行全称（含支行）
	BankBranchId  string `json:"bank_branch_id"` // 开户银行联行号
	AccountNumber string `json:"account_number"` // 银行账号
	VerifyResult  string `json:"verify_result"`  // 汇款验证结果
}

// 提交申请单API
// doc: https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/tool/applyment4sub/chapter3_1.shtml
func Applyment(cli *core.Client, params *ApplymentReq) (applymentId string, err error) {
	var body, resp []byte
	var reply *ApplymentReply

	body, _ = json.Marshal(parmas)
	resp, err = cli.DoPost("/v3/applyment4sub/applyment/", string(body), true)
	if err != nil {
		return
	}
	reply = new(ApplymentReply)

	err = json.Unmarshal(resp, reply)
	if err != nil {
		return
	}

	applymentId = reply.ApplymentId
	return
}

// 查询申请单状态API
// doc: https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/tool/applyment4sub/chapter3_2.shtml
func QueryApplymentState(cli *core.Client, applymentId string) (state *ApplymentQueryReply, err error) {
	var resp []byte

	resp, err = cli.DoGet(fmt.Sprintf("/v3/applyment4sub/applyment/applyment_id/%s", applymentId), true)
	if err != nil {
		return
	}

	state = new(ApplymentQueryReply)
	err = json.Unmarshal(resp, state)
	return
}

// 通过业务申请编号查询申请状态
func QueryApplymentStateByCode(cli *core.Client, bussinessCode string) (state *ApplymentQueryReply, err error) {
	var resp []byte

	resp, err = cli.DoGet(fmt.Sprintf("/v3/applyment4sub/applyment/business_code/%s", bussinessCode), true)
	if err != nil {
		return
	}

	state = new(ApplymentQueryReply)
	err = json.Unmarshal(resp, state)
	return
}

// 修改结算帐号API
// doc: https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/tool/applyment4sub/chapter3_3.shtml
func ModifySettlement(cli *core.Client, params *ModifySettlementReq) (flag bool, err error) {
	var body []byte
	var code int

	body, _ = json.Marshal(parmas)
	_, code, err = cli.DoPost(fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/modify-settlement", params.SubMchid), string(body))
	if err != nil {
		return
	}
	if code == 204 {
		flag = true
	}
	return
}

// 查询结算账户API
func QuerySettlement(cli *core.Client, subMchid string) (reply *QuerySettlementReply, err error) {
	var resp []byte

	resp, err = cli.DoGet(fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/settlement", subMchid), true)
	if err != nil {
		return
	}

	reply = new(QuerySettlementReply)
	err = json.Unmarshal(resp, reply)
	return
}
