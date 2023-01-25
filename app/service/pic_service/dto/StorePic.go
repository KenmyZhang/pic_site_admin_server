package dto

type StorePic struct {
	Id         int64    `json:"id"`
	ListPicUrl []string `json:"list_pic_url" valid:"Required;"`
	Mobile     string   `json:"mobile"`
	Wechat     string   `json:"wechat"`
	PicType    int      `json:"pic_type"`
}
