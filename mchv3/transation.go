package mchv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	PayTypeAPP    = "app"
	PayTypeJSAPI  = "jsapi"
	PayTypeNative = "native"
	PayTypeH5     = "h5"
)

type SceneInfo struct {
	ClientIP  string     `json:"payer_client_ip"`
	DeviceID  string     `json:"device_id,omitempty"`
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

type StoreInfo struct {
	ID       string `json:"id"` // 公众号ID
	Name     string `json:"name,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	Address  string `json:"address,omitempty"`
}

type PayDetail struct {
	CostPrice *int64       `json:"cost_price,omitempty"`
	InvoiceID string       `json:"invoice_id,omitempty"`
	Goods     []*GoodsInfo `json:"goods_detail"`
}

type GoodsInfo struct {
	ID        string `json:"merchant_goods_id"` // 公众号ID
	WechatID  string `json:"wechatpay_goods_id,omitempty"`
	GoodsName string `json:"goods_name,omitempty"`
	Quantity  int64  `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
}

type SettleInfo struct {
	ProfitSharing bool   `json:"profit_sharing"`
	SubsidyAmount *int64 `json:"subsidy_amount,omitempty"`
}

type PayOrder struct {
	Type string `json:"-"`
	// isv模式，除了sub app 均为必须
	SpAppID  string `json:"sp_appid,omitempty"`  // 服务商appid
	SpMchID  string `json:"sp_mchid,omitempty"`  // 服务商mchid
	SubAppID string `json:"sub_appid,omitempty"` // 子账户appid
	SubMchID string `json:"sub_mchid,omitempty"` // 子商户mchid
	// 非isv模式，必须
	AppID       string     `json:"appid,omitempty"` // 公众号ID
	MchID       string     `json:"mchid,omitempty"`
	Description string     `json:"description"`
	OutTradeNo  string     `json:"out_trade_no"`
	TimeExpire  *time.Time `json:"time_expire,omitempty"`
	Attach      string     `json:"attach,omitempty"`
	NotifyURL   string     `json:"notify_url"`
	GoodsTag    string     `json:"goods_tag,omitempty"`
	Amount      struct {
		Total    int64  `json:"total"`
		Currency string `json:"currency,omitempty"`
	} `json:"amount"`
	Payer struct {
		OpenID    string `json:"openid,omitempty"`
		SpOpenID  string `json:"sp_openid,omitempty"`
		SubOpenID string `json:"sub_openid,omitempty"`
	} `json:"payer"`
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	PayDetail  *PayDetail  `json:"detail,omitempty"`
	SceneInfo  *SceneInfo  `json:"scene_info,omitempty"`
}

type CombineSubOrder struct {
	MchID       string `json:"mchid"`
	SubMchID    string `json:"sub_mchid"`
	Description string `json:"description"`
	OutTradeNo  string `json:"out_trade_no"`
	Attach      string `json:"attach"`
	Amount      struct {
		Total    int64  `json:"total"`
		Currency string `json:"currency,omitempty"`
	} `json:"amount"`
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

type CombinePayOrder struct {
	Type       string             `json:"-"`
	AppID      string             `json:"combine_appid,omitempty"` // 服务商appid
	MchID      string             `json:"combine_mchid,omitempty"` // 服务商mchid
	OutTradeNo string             `json:"combine_out_trade_no"`
	SceneInfo  *SceneInfo         `json:"scene_info,omitempty"`
	SubOrders  []*CombineSubOrder `json:"sub_orders"`
	PayInfo    struct {
		OpenID string `json:"openid,omitempty"`
	} `json:"combine_payer_info,omitempty"`
	TimeStart  *time.Time `json:"time_start,omitempty"`
	TimeExpire *time.Time `json:"time_expire,omitempty"`
	NotifyURL  string     `json:"notify_url"`
}

func Transaction(v3 *Client, order PayOrder) (result map[string]string, err error) {
	partner := ""
	if v3.Isv {
		if order.SpAppID == "" || order.SpMchID == "" || order.SubMchID == "" {
			err = errors.New("商户号缺失，无法下单")
			return
		}
		order.AppID = ""
		order.MchID = ""
		partner = "partner/"
	} else {
		if order.AppID == "" || order.MchID == "" {
			err = errors.New("商户号缺失，无法下单")
			return
		}
		order.SpAppID = ""
		order.SpMchID = ""
		order.SubMchID = ""
		order.SubAppID = ""
	}
	url := fmt.Sprintf("/v3/pay/%stransactions/%s", partner, order.Type)
	data, _ := json.Marshal(order)
	body, _, err := v3.DoPost(url, string(data))
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &result)
	return
}

func CombineTransaction(v3 *Client, order CombinePayOrder) (result map[string]string, err error) {
	url := fmt.Sprintf("/v3/combine-transactions/%s", order.Type)
	data, _ := json.Marshal(order)
	body, _, err := v3.DoPost(url, string(data))
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &result)
	return
}
