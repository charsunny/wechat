package menu

// DatabaseTable

//  商品
type Product struct {
	ProductBase *ProductBase `json:"product_base"`
	SkuList []*SkuItem	`json:"sku_list"`
	AttrExt *AttrExt 	`json:"attrext"`
	DeliveryInfo *DeliveryInfo `json:"delivery_info"`
}

type ProductBase struct {
	CategoryId []string                     `json:"category_id"`
	Property   []*Property          `json:"property"`
	Name       string                       `json:"name"`
	SkuInfo    []*SkuInfo           `json:"sku_info"`
	MainImg    string                       `json:"main_img"`
	Img        []string                     `json:"img"`
	Detail     []*ProductBaseDetail `json:"detail"`
	BuyLimit   int                          `json:"buy_limit"`
}

// SkuItem 库存
type SkuItem struct {
	SkuId       string `json:"sku_id"`
	Price       int64  `json:"price"`
	IconUrl     string `json:"icon_url"`
	ProductCode string `json:"product_code"`
	OriPrice    int64  `json:"ori_price"`
	Quantity    int    `json:"quantity"`
}

// Express 快递模板
type Express struct {
	TemplateId int    `json:"template_id"`
	Name       string `json:"name"`
	Assumer    int    `json:"assumer"`
	Valuation  int    `json:"valuation"`
	TopFee     string `json:"top_fee"`
	CreateTime string
}

// API Struct
type ExpressTemplate struct {
	Name      string       `json:"name"`
	Assumer   int          `json:"assumer"`
	Valuation int          `json:"valuation"`
	TopFee    []TopFeeItem `json:"top_fee"`
}

type TopFeeItem struct {
	Type   int `json:"Type"`
	Normal struct {
		StartStandards int `json:"StartStandards"`
		StartFees      int `json:"StartFees"`
		AddStandards   int `json:"AddStandards"`
		AddFees        int `json:"AddFees"`
	} `json:"Normal"`
	Custom []struct {
		StartStandards int    `json:"StartStandards"`
		StartFees      int    `json:"StartFees"`
		AddStandards   int    `json:"AddStandards"`
		AddFees        int    `json:"AddFees"`
		DestCountry    string `json:"DestCountry"`
		DestProvince   string `json:"DestProvince"`
		DestCity       string `json:"DestCity"`
	} `json:"Custom"`
}

type AddExpressTemplateRequest struct {
	ExpressTemplate ExpressTemplate `json:"delivery_template"`
}

type AddExpressTemplateResponse struct {
	TemplateId string `json:"template_id"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type Property struct {
	Id  string `json:"id"`
	Vid string `json:"vid"`
}

type SkuInfo struct {
	Id  string   `json:"id"`
	Vid []string `json:"vid"`
}

type ProductBaseDetail struct {
	Text string `json:"text"`
	Img  string `json:"img`
}

type SkuListItem struct {
	SkuId       string `json:"sku_id"`
	Price       int64  `json:"price"`
	IconUrl     string `json:"icon_url"`
	ProductCode string `json:"product_code"`
	OriPrice    int64  `json:"ori_price"`
	Quantity    int    `json:"quantity"`
}

type AttrExtLocation struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
}

type AttrExt struct {
	Location         *AttrExtLocation `json:"location"`
	IsPostFree       int                      `json:"isPostFree"`
	IsHasReceipt     int                      `json:"isHasReceipt"`
	IsUnderGuaranty  int                      `json:"isUnderGuaranty"`
	IsSupportReplace int                      `json:"isSupportReplace"`
}

type DeliveryInfo struct {
	DeliveryType int                            `json:"delivery_type"`
	TemplateId   int                            `json:"template_id"`
	Express      []*DeliveryInfoExpress `json:"express"`
}

type DeliveryInfoExpress struct {
	Id    int64 `json:"id"`
	Price int64 `json:"price"`
}

type AddRequest struct {
	ProductBase  ProductBase   `json:"product_base"`
	SkuList      []SkuListItem `json:"sku_list"`
	AttrExt      AttrExt       `json:"attrext"`
	DeliveryInfo *DeliveryInfo  `json:"delivery_info"`
}

type AddResponse struct {
	ProductId string `json:"product_id"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type DeleteRequest struct {
	ProductId string `json:"product_id"`
}

type UpdateRequest struct {
	ProductId    string                 `json:"product_id"`
	ProductBase  ProductBase   `json:"product_base"`
	SkuList      []*SkuListItem `json:"sku_list"`
	AttrExt      AttrExt       `json:"attrext"`
	DeliveryInfo *DeliveryInfo  `json:"delivery_info"`
}

type Detail struct {
	ProductId    string                 `json:"product_id"`
	ProductBase  ProductBase   `json:"product_base"`
	SkuList      []*SkuListItem `json:"sku_list"`
	AttrExt      AttrExt       `json:"attrext"`
	DeliveryInfo *DeliveryInfo  `json:"delivery_info"`
}

type GetRequest struct {
	ProductId string `json:"product_id"`
}

type GetResponse struct {
	ProductInfo Detail `json:"product_info"`
	ErrCode     int             `json:"errcode"`
	ErrMsg      string          `json:"errmsg"`
}

type GetByStatusRequest struct {
	Status int `json:"status"`
}

type GetByStatusResponse struct {
	ProductInfo []Detail `json:"product_info"`
	ErrCode     int               `json:"errcode"`
	ErrMsg      string            `json:"errmsg"`
}

type UpdateStatusRequest struct {
	ProductId       string `json:"product_id"`
	Status int    `json:"status"`
}

type GetSubClassesByClassifyRequest struct {
	CateId int64 `json:"cate_id"`
}

type CateInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetSubClassesByClassifyResponse struct {
	CateList []CateInfo `json:"cate_list"`
	ErrCode  int                 `json:"errcode"`
	ErrMsg   string              `json:"errmsg"`
}

type GetAllSkuByClassifyRequest struct {
	CateId string `json:"cate_id"`
}

type SkuValueList struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SkuTable struct {
	Id        string                  `json:"id"`
	Name      string                  `json:"name"`
	ValueList []SkuValueList `json:"value_list"`
}

type GetAllSkuByClassifyResponse struct {
	SkuTable SkuTable `json:"sku_table"`
	ErrCode  int               `json:"errcode"`
	ErrMsg   string            `json:"errmsg"`
}

type GetAllPropertyByClassifyRequest struct {
	CateId string `json:"cate_id"`
}

type ClassifyPropertyValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ClassifyProperty struct {
	Id            string                           `json:"id"`
	Name          string                           `json:"name"`
	PropertyValue []ClassifyPropertyValue `json:"property_value"`
}

type GetAllPropertyByClassifyResponse struct {
	Properties ClassifyProperty `json:"properties"`
	ErrCode    int                       `json:"errcode"`
	ErrMsg     string                    `json:"errmsg"`
}

type StockSkuInfo struct {
	Id  string
	Vid string
}

type AddStockRequest struct {
	ProductId string
	SkuInfo   []StockSkuInfo
	Quantity  int
}

type CReduceStockRequest struct {
	ProductId string
	SkuInfo   []StockSkuInfo
	Quantity  int
}