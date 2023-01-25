package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type NideshopAd struct {
	AdPositionId int                   `json:"ad_position_id"`
	Content      string                `json:"content"`
	Enabled      int                   `json:"enabled"`
	EndTime      int                   `json:"end_time"`
	Id           int                   `json:"id"`
	ImageUrl     string                `json:"image_url"`
	Link         string                `json:"link"`
	MediaType    int                   `json:"media_type"`
	Name         string                `json:"name"`
	UpdateTime   time.Time             `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime   time.Time             `json:"createTime" gorm:"autoCreateTime"`
	IsDel        soft_delete.DeletedAt `json:"isDel" gorm:"softDelete:flag"`
}

func (NideshopAd) TableName() string {
	return "nideshop_ad"
}

// get all
func GetAllNideshopAds(pageNUm int, pageSize int, maps interface{}) (int64, []NideshopAd) {
	var (
		total int64
		data  []NideshopAd
	)

	db.Model(&NideshopAd{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddNideshopAd(m *NideshopAd) error {
	var err error
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateNideshopAd(m *NideshopAd) error {
	var err error
	m.UpdateTime = time.Now()
	oldTopic := &NideshopAd{}
	db.Model(&NideshopAd{}).Where("id = ?", m.Id).First(oldTopic)
	m.CreateTime = oldTopic.CreateTime
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func GetNideShopAd(id int64) *NideshopAd {
	shopAd := &NideshopAd{}
	db.Model(&NideshopAd{}).Where("id = ?", id).First(shopAd)
	return shopAd
}

func DelNideshopAd(ids []int64) error {
	err := db.Where("id in (?)", ids).Delete(&NideshopAd{}).Error
	if err != nil {
		return err
	}

	return err
}
