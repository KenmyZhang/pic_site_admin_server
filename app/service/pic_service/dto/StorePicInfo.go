package dto

type StorePicInfo struct {
	Id            int64    `json:"id"`
	PrimaryPicUrl string   `json:"primary_pic_url" valid:"Required;"`
	ListPicUrl    []string `json:"list_pic_url" valid:"Required;"`
	Mobile        string   `json:"mobile"`
	Wechat        string   `json:"wechat"`
	PicType       int      `json:"pic_type"`
}
