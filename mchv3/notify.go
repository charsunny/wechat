package mchv3

import "encoding/json"

type PayResp struct {
	ID           string      `json:"id"`
	CreateTime   Time        `json:"create_time"`
	EventType    string      `json:"event_type"`
	ResourceType string      `json:"resource_type"`
	Resource     EncryptData `json:"resource"`
	Summary      string      `json:"summary"`
}

type PayResult struct {
	Type string `json:"-"`
	// isv模式，除了sub app 均为必须
	SpAppID  string `json:"sp_appid,omitempty"`  // 服务商appid
	SpMchID  string `json:"sp_mchid,omitempty"`  // 服务商mchid
	SubAppID string `json:"sub_appid,omitempty"` // 子账户appid
	SubMchID string `json:"sub_mchid,omitempty"` // 子商户mchid
	// 合单支付模式
	CombineAppID      string              `json:"combine_appid,omitempty"` // 公众号ID
	CombineMchID      string              `json:"combine_mchid,omitempty"`
	CombineOutTradeNo string              `json:"combine_out_trade_no"`
	CombinePayResult  []*CombinePayResult `json:"sub_orders"`
	CombinePayInfo    struct {
		OpenID string `json:"openid,omitempty"`
	} `json:"combine_payer_info,omitempty"`
	// 非isv模式，必须
	AppID      string `json:"appid,omitempty"` // 公众号ID
	MchID      string `json:"mchid,omitempty"`
	OutTradeNo string `json:"out_trade_no"`

	TransactionID  string `json:"transaction_id"`
	TradeType      string `json:"trade_type"`                 // 交易类型
	TradeState     string `json:"trade_state"`                // 交易类型
	TradeStateDesc string `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType       string `json:"bank_type,omitempty"`        // 银行类型
	Attach         string `json:"attach,omitempty"`
	SuccessTime    *Time  `json:"success_time,omitempty"`
	Amount         struct {
		Total         int64  `json:"total"`
		PayerTotal    int64  `json:"payer_total"`
		Currency      string `json:"currency,omitempty"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
	Payer struct {
		OpenID    string `json:"openid,omitempty"`
		SpOpenID  string `json:"sp_openid,omitempty"`
		SubOpenID string `json:"sub_openid,omitempty"`
	} `json:"payer"`
	PromotionDetail *PromotionDetail `json:"promotion_detail,omitempty"`
	SceneInfo       struct {
		DeviceID string `json:"device_id"`
	} `json:"scene_info,omitempty"`
}

type CombinePayResult struct {
	MchID          string `json:"mchid,omitempty"`     // 服务商mchid
	SubMchID       string `json:"sub_mchid,omitempty"` // 子商户mchid
	TransactionID  string `json:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no"`
	TradeType      string `json:"trade_type"`                 // 交易类型
	TradeState     string `json:"trade_state"`                // 交易类型
	TradeStateDesc string `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType       string `json:"bank_type,omitempty"`        // 银行类型
	Attach         string `json:"attach,omitempty"`
	SuccessTime    *Time  `json:"success_time,omitempty"`
	Amount         struct {
		Total         int64  `json:"total"`
		PayerTotal    int64  `json:"payer_total"`
		Currency      string `json:"currency,omitempty"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}

func (cli *Client) GetPayResult(body []byte) (resp PayResp, result PayResult, err error) {
	if err = json.Unmarshal(body, &resp); err != nil {
		return
	}
	data, err := cli.DecryptData(resp.Resource)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return
	}
	return
}
