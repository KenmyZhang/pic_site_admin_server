package casbin

import (
	"log"

	"github.com/KenmyZhang/pic_site_admin_server/pkg/constant"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/runtime"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Printf("[info] casbin %s", err)
	}
	model, err := model.NewModelFromFile("conf/rbac_model.conf")
	e, err := casbin.NewSyncedEnforcer(model, a)
	if err != nil {
		log.Printf("[info] casbin %s", err)
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Printf("[info] casbin %s", err)
	}

	runtime.Runtime.SetCasbin(constant.YSHOP_CASBIN, e)
}
