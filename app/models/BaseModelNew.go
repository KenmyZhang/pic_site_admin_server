package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModelNew struct {
	Id         int64                 `gorm:"primary_key;autoIncrement" json:"id"` // 主键id
	UpdateTime *time.Time            `json:"updateTime" gorm:"autoUpdateTime"`    // 更新时间
	CreateTime *time.Time            `json:"createTime" gorm:"autoCreateTime"`    // 创建时间
	IsDel      soft_delete.DeletedAt `json:"isDel" gorm:"softDelete:flag"`        // 删除标签 0 未删除 1 已删除
}
