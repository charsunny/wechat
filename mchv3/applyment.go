package mchv3

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PayApply struct {
	BusinessCode string      `json:"business_code"`
	ContactInfo  ContactInfo `json:"contact_info"`
	SubjectInfo  SubjectInfo `json:"subject_info"`
}

type ContactInfo struct {
	ContactName     string `json:"contact_name"`
	ContactIDNumber string `json:"contact_id_number,omitempty"`
	OpenID          string `json:"open_id,omitempty"`
	MobilePhone     string `json:"mobile_phone"`
	ContactEmail    string `json:"contact_email"`
}

type SubjectInfo struct {
	SubjectType           string               `json:"subject_type"`
	BusinessLicenseInfo   *BusinessLicenseInfo `json:"business_license_info,omitempty"`
	CertificateInfo       *CertificateInfo     `json:"certificate_info,omitempty"`
	OrganizationInfo      *OrganizationInfo    `json:"organization_info,omitempty"`
	CertificateLetterCopy string               `json:"certificate_letter_copy,omitempty"`
	IdentityInfo          IdentityInfo         `json:"identity_info"`
	BusinessInfo          BusinessInfo         `json:"business_info"`
	SettlementInfo        SettlementInfo       `json:"settlement_info"`
	BankAccountInfo       *BankAccountInfo     `json:"bank_account_info,omitempty"`
	AdditionInfo          *AdditionInfo        `json:"addition_info,omitempty"`
}

type BusinessLicenseInfo struct {
	LicenseCopy   string `json:"license_copy"`
	LicenseNumber string `json:"license_number"`
	MerchantName  string `json:"merchant_name"`
	LegalPerson   string `json:"legal_person"`
}

type CertificateInfo struct {
	CertCopy       string `json:"cert_copy"`
	CertType       string `json:"cert_type"`
	CertNumber     string `json:"cert_number"`
	MerchantName   string `json:"merchant_name"`
	CompanyAddress string `json:"company_address"`
	LegalPerson    string `json:"legal_person"`
	PeriodBegin    string `json:"period_begin"`
	PeriodEnd      string `json:"period_end"`
}

type OrganizationInfo struct {
	OrganizationCopy string `json:"organization_copy"`
	OrganizationCode string `json:"organization_code"`
	OrgPeriodBegin   string `json:"org_period_begin"`
	OrgPeriodEnd     string `json:"org_period_end"`
}

type IdentityInfo struct {
	IDDocType  string     `json:"id_doc_type"`
	IDCardInfo IDCardInfo `json:"id_card_info"`
	Owner      bool       `json:"owner"`
}

type IDCardInfo struct {
	IDCardCopy      string `json:"id_card_copy"`
	IDCardNational  string `json:"id_card_national"`
	IDCardName      string `json:"id_card_name"`
	IDCardNumber    string `json:"id_card_number"`
	CardPeriodBegin string `json:"card_period_begin"`
	CardPeriodEnd   string `json:"card_period_end"`
}

type BusinessInfo struct {
	MerchantShortname string    `json:"merchant_shortname"`
	ServicePhone      string    `json:"service_phone"`
	SalesInfo         SalesInfo `json:"sales_info"`
}

type SalesInfo struct {
	SalesScenesType string           `json:"sales_scenes_type"`
	BizStoreInfo    *BizStoreInfo    `json:"biz_store_info,omitempty"`
	MpInfo          *MpInfo          `json:"mp_info,omitempty"`
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
	AppInfo         *AppInfo         `json:"app_info,omitempty"`
}

type BizStoreInfo struct {
	BizStoreName     string `json:"biz_store_name"`
	BizAddressCode   string `json:"biz_address_code"`
	BizStoreAddress  string `json:"biz_store_address"`
	StoreEntrancePic string `json:"store_entrance_pic"`
	IndoorPic        string `json:"indoor_pic"`
	BizSubAppid      string `json:"biz_sub_appid"`
}

type MpInfo struct {
	MpAppid    string   `json:"mp_appid,omitempty"`
	MpSubAppid string   `json:"mp_sub_appid,omitempty"`
	MpPics     []string `json:"mp_pics"`
}

type MiniProgramInfo struct {
	MiniProgramAppid    string   `json:"mini_program_appid,omitempty"`
	MiniProgramSubAppid string   `json:"mini_program_sub_appid,omitempty"`
	MiniProgramPics     []string `json:"mini_program_pics"`
}

type AppInfo struct {
	AppAppid    string   `json:"app_appid,omitempty"`
	AppSubAppid string   `json:"app_sub_appid,omitempty"`
	AppPics     []string `json:"app_pics"`
}

type SettlementInfo struct {
	SettlementID        string   `json:"settlement_id"`
	QualificationType   string   `json:"qualification_type"`
	Qualifications      []string `json:"qualifications,omitempty"`
	ActivitiesID        string   `json:"activities_id,omitempty"`
	ActivitiesRate      string   `json:"activities_rate,omitempty"`
	ActivitiesAdditions []string `json:"activities_additions,omitempty"`
}

type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"`
	AccountName     string `json:"account_name"`
	AccountBank     string `json:"account_bank"`
	BankAddressCode string `json:"bank_address_code"`
	BankBranchID    string `json:"bank_branch_id,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	AccountNumber   string `json:"account_number"`
}

type AdditionInfo struct {
	LegalPersonCommitment string   `json:"legal_person_commitment,omitempty"`
	LegalPersonVideo      string   `json:"legal_person_video,omitempty"`
	BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"`
	BusinessAdditionMsg   string   `json:"business_addition_msg,omitempty"`
}

// 商户进件接口
func Applyment(v3 *Client, apply PayApply) (applyment_id string, err error) {
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/applyment4sub/applyment")
	data, _ := json.Marshal(apply)
	var result struct {
		ApplymentID string `json:"applyment_id"`
	}
	body, err := v3.DoPost(url, string(data))
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &result)
	applyment_id = result.ApplymentID
	if applyment_id == "" {
		err = errors.New(string(body))
	}
	return
}

type ApplymentResult struct {
	BusinessCode      string `json:"business_code"`
	ApplymentID       string `json:"applyment_id"`
	SubMchID          string `json:"sub_mchid,omitempty"`
	SignURL           string `json:"sign_url,omitempty"`
	ApplymentState    string `json:"applyment_state"`
	ApplymentStateMsg string `json:"applyment_state_msg"`
	AuditDetail       struct {
		Field        string `json:"field"`
		FieldName    string `json:"field_name"`
		RejectReason string `json:"reject_reason"`
	} `json:"audit_detail,omitempty"`
}

// 商户进件接口
func QueryApplymentByCode(v3 *Client, business_code string) (result ApplymentResult, err error) {
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/applyment4sub/applyment/business_code/%s", business_code)
	body, err := v3.DoGet(url)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &result)
	return
}

// 商户进件接口
func QueryApplymentByID(v3 *Client, applyment_id string) (result ApplymentResult, err error) {
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/applyment4sub/applyment/applyment_id/%s", applyment_id)
	body, err := v3.DoGet(url)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &result)
	return
}

type SettlementModify struct {
	SubMchID        string `json:"sub_mchid"`
	BankAccountType string `json:"account_type"`
	AccountBank     string `json:"account_bank"`
	BankAddressCode string `json:"bank_address_code"`
	BankBranchID    string `json:"bank_branch_id,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	AccountNumber   string `json:"account_number"`
}

func ModifySettlement(v3 *Client, modify SettlementModify) (err error) {
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/apply4sub/sub_merchants/%s/modify-settlement", modify.SubMchID)
	data, _ := json.Marshal(modify)

	_, err = v3.DoPost(url, string(data))
	return
}

type SettlementResult struct {
	VerifyResult    string `json:"verify_result"`
	BankAccountType string `json:"account_type"`
	AccountBank     string `json:"account_bank"`
	BankAddressCode string `json:"bank_address_code"`
	BankBranchID    string `json:"bank_branch_id,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	AccountNumber   string `json:"account_number"`
}

func QuerySettlement(v3 *Client, mchid string) (result SettlementModify, err error) {
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/apply4sub/sub_merchants/%s/settlement", mchid)
	body, err := v3.DoGet(url)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &result)
	return
}
