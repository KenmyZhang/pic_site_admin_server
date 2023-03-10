package global

import (
	"github.com/KenmyZhang/pic_site_admin_server/conf"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	YSHOP_DB               *gorm.DB
	YSHOP_VP               *viper.Viper
	YSHOP_LOG              *zap.SugaredLogger
	YSHOP_CONFIG           conf.Config
	YSHOP_OFFICIAL_ACCOUNT *officialaccount.OfficialAccount
)
