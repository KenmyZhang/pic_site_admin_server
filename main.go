package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KenmyZhang/pic_site_admin_server/app/listen"
	"github.com/KenmyZhang/pic_site_admin_server/app/models"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/base"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/global"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/logging"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/redis"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/wechat"
	"github.com/KenmyZhang/pic_site_admin_server/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	global.YSHOP_VP = base.Viper()
	global.YSHOP_LOG = base.SetupLogger()
	models.Setup()
	logging.Setup()
	redis.Setup()
	listen.Setup()
	wechat.InitWechat()
}

// @title gin-shop  API
// @version 1.0.1
// @description gin-shop商城后台管理系统
// @termsOfService https://gitee.com/guchengwuyue/gin-shop
// @license.name apache2
func main() {
	gin.SetMode(global.YSHOP_CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.YSHOP_CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	global.YSHOP_LOG.Info("[info] start http server listening %s", endPoint)
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
