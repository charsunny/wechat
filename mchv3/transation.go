package mchv3

import (
	"encoding/json"
	"errors"
	"fmt"
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

type PromotionDetail struct {
	CouponID            string         `json:"coupon_id"`
	Name                string         `json:"name"`
	Scope               string         `json:"scope"` // GLOBAL：全场代金券 SINGLE：单品优惠
	Type                string         `json:"type"`
	Amount              int64          `json:"amount"`
	StockID             string         `json:"stock_id"`
	WechatpayContribute int64          `json:"wechatpay_contribute"`
	MerchantContribute  int64          `json:"merchant_contribute"`
	OtherContribute     int64          `json:"other_contribute"`
	Currency            string         `json:"currency,omitempty"`
	GoodsDetails        []*GoodsDetail `json:"goods_detail,omitempty"`
}

type GoodsDetail struct {
	GoodsID        string `json:"goods_id"` // 公众号ID
	Quantity       int64  `json:"quantity"`
	UnitPrice      int64  `json:"unit_price"`
	DiscountAmount int64  `json:"discount_amount,omitempty"`
	GoodsRemark    string `json:"goods_remark,omitempty"`
}

type GoodsInfo struct {
	ID        string `json:"merchant_goods_id"` // 小店商品id
	WechatID  string `json:"wechatpay_goods_id,omitempty"`
	GoodsName string `json:"goods_name,omitempty"`
	Quantity  int64  `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
}

// 用于电商收付通支付
type SettleInfo struct {
	ProfitSharing bool   `json:"profit_sharing"`           // 是否分账
	SubsidyAmount *int64 `json:"subsidy_amount,omitempty"` // 分账补差金额
}

type PayOrder struct {
	Type string `json:"-"`
	// isv模式，除了sub app 均为必须
	SpAppID  string `json:"sp_appid,omitempty"`  // 服务商appid
	SpMchID  string `json:"sp_mchid,omitempty"`  // 服务商mchid
	SubAppID string `json:"sub_appid,omitempty"` // 子账户appid
	SubMchID string `json:"sub_mchid,omitempty"` // 子商户mchid
	// 非isv模式，必须
	AppID         string `json:"appid,omitempty"` // 公众号ID
	MchID         string `json:"mchid,omitempty"`
	Description   string `json:"description"`
	OutTradeNo    string `json:"out_trade_no"`
	TimeExpire    *Time  `json:"time_expire,omitempty"`
	Attach        string `json:"attach,omitempty"`
	NotifyURL     string `json:"notify_url"`
	GoodsTag      string `json:"goods_tag,omitempty"`
	ProfitSharing bool   `json:"profit_sharing"` // 是否分账
	Amount        struct {
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
	TimeStart  *Time  `json:"time_start,omitempty"`
	TimeExpire *Time  `json:"time_expire,omitempty"`
	NotifyURL  string `json:"notify_url"`
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
