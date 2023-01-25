package models

import "time"

type YshopStorePic struct {
	Id            int64     `json:"id"`
	ListPicUrl    string    `json:"list_pic_url"`    // 商品列表图
	PrimaryPicUrl string    `json:"primary_pic_url"` // 主图
	Mobile        string    `json:"mobile"`
	Wechat        string    `json:"wechat"`
	PicType       int       `json:"pic_type"`
	UpdateTime    time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime    time.Time `json:"createTime" gorm:"autoCreateTime"`
}

func (YshopStorePic) TableName() string {
	//return "yshop_store_product"
	return "nideshop_goods"
}

func GetProduct(id int64) YshopStorePic {
	var product YshopStorePic
	db.Where("id =  ?", id).First(&product)

	return product
}

// get all
func GetAllProduct(pageNUm int, pageSize int, maps interface{}) (int64, []YshopStorePic) {
	var (
		total int64
		data  []YshopStorePic
	)

	db.Model(&YshopStorePic{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("ProductCate").Order("id desc").Find(&data)

	return total, data
}

func AddProduct(m *YshopStorePic) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByProduct(id int64, m YshopStorePic) error {
	var err error
	err = db.Model(&YshopStorePic{}).Where("id = ?", id).UpdateColumns(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByProduct(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&YshopStorePic{}).Error
	if err != nil {
		return err
	}

	return err
}
