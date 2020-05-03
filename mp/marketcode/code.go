package marketcode

type MarketCodeApply struct {
	Status string `json:"status"` // INIT PROCESSING FINISH为可下载
	IsvApplicationId string `json:"isv_application_id"` //
	ApplicationId int64 `json:"application_id"` //
	CreateTime string `json:"create_time"` //
	UpdateTime string `json:"update_time"` //
	CodeGenerateList []struct{
		CodeStart int64 `json:"code_start"` // 包含，如0
		CodeEnd int64 `json:"code_end"` // 包含，如49999，上述0-49999为一个号码段
	} `json:"code_generate_list"`
}

type ActiveCodeInfo struct {
	ApplicationId int64 `json:"application_id"` //
	ActivityName string `json:"activity_name"` // 数据分析活动区分依据，请规范命名
	ProductBrand string `json:"product_brand"` // 数据分析品牌区分依据，请规范命名
	ProductTitle string `json:"product_title"` // 数据分析商品区分依据，请规范命名
	ProductCode string `json:"product_code"` // EAN商品条码，请规范填写
	WxaAppid string `json:"wxa_appid"` 		//	扫码跳转小程序的appid
	WxaPath string `json:"wxa_path"` 		//	扫码跳转小程序的appid
	WxaType int `json:"wxa_type"`			// 默认为0正式版，开发版为1，体验版为2
	CodeStart int64 `json:"code_start"` // 包含，如0
	CodeEnd int64 `json:"code_end"` // 包含，如49999，上述0-49999为一个号码段
}
