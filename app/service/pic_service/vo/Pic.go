package vo

type Product struct {
	Id            int64    `json:"id"`
	Image         string   `json:"image"`
	ListPicUrl    string   `json:"list_pic_url"`
	ListPicUrlArr []string `json:"list_pic_url_arr"`
	StoreName     string   `json:"storeName"`
	StoreInfo     string   `json:"storeInfo"`
	Keyword       string   `json:"keyword"`
	CategoryId    int64    `json:"category_id" valid:"Required;"`
	Price         float64  `json:"price"`
	VipPrice      float64  `json:"vipPrice"`
	OtPrice       float64  `json:"otPrice"`
	Postage       float64  `json:"postage"`
	GoodsUnit     string   `json:"goods_unit"`
	Sort          int16    `json:"sort"`
	Sales         int      `json:"sales"`
	Stock         int      `json:"stock"`
	Description   string   `json:"description"`
	IsPostage     int8     `json:"isPostage"`
	GiveIntegral  int      `json:"giveIntegral"`
	Cost          float64  `json:"cost"`
	IsGood        int8     `json:"isGood"`
	Ficti         int      `json:"ficti"`
	Browse        int      `json:"browse"`
	IsSub         int8     `json:"isSub"`
	TempId        int64    `json:"tempId"`
	SpecType      int8     `json:"specType"`
	IsIntegral    int8     `json:"isIntegral"`
	Integral      int32    `json:"integral"`
	UserCollect   bool     `json:"userCollect"`
}
