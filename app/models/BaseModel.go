package models

import (
	"fmt"

	"github.com/KenmyZhang/pic_site_admin_server/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/soft_delete"

	//"gorm.io/plugin/soft_delete"
	"log"
	"os"
	"time"

	"github.com/KenmyZhang/pic_site_admin_server/pkg/casbin"
)

var db *gorm.DB

type BaseModel struct {
	Id         int64                 `gorm:"primary_key;autoIncrement" json:"id"` // 主键id
	UpdateTime time.Time             `json:"updateTime" gorm:"autoUpdateTime"`    // 更新时间
	CreateTime time.Time             `json:"createTime" gorm:"autoCreateTime"`    // 创建时间
	IsDel      soft_delete.DeletedAt `json:"isDel" gorm:"softDelete:flag"`        // 删除标签 0 未删除 1 已删除
}

// Setup initializes the database instance
func Setup() {
	var err error
	var connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.YSHOP_CONFIG.Database.User,
		global.YSHOP_CONFIG.Database.Password,
		global.YSHOP_CONFIG.Database.Host,
		global.YSHOP_CONFIG.Database.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Printf("[info] gorm %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("[info] gorm %s", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.YSHOP_DB = db

	casbin.InitCasbin(db)

}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
