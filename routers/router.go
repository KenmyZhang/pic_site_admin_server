package routers

import (
	"net/http"

	"github.com/KenmyZhang/pic_site_admin_server/app/controllers/admin"
	_ "github.com/KenmyZhang/pic_site_admin_server/docs"
	"github.com/KenmyZhang/pic_site_admin_server/middleware"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/upload"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	loginController := admin.LoginController{}
	r.POST("/auth/login", loginController.Login)
	r.GET("/auth/captcha", loginController.Captcha)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	menuController := admin.MenuController{}
	userController := admin.UserController{}
	deptController := admin.DeptController{}
	dictDetailController := admin.DictDetailController{}
	dictController := admin.DictController{}
	roleController := admin.RoleController{}
	logController := admin.LogController{}
	materialController := admin.MaterialController{}
	materialGroupController := admin.MaterialGroupController{}
	adminRouter := r.Group("/admin")
	adminRouter.Use(middleware.Jwt()).Use(middleware.Log())

	{
		adminRouter.GET("/auth/info", loginController.Info)
		adminRouter.DELETE("/auth/logout", loginController.Logout)

		adminRouter.GET("/material", materialController.GetAll)
		adminRouter.POST("/material", materialController.Post)
		adminRouter.PUT("/material", materialController.Put)
		adminRouter.DELETE("/material/:id", materialController.Delete)
		adminRouter.POST("/material/upload", materialController.Upload)

		adminRouter.GET("/materialgroup", materialGroupController.GetAll)
		adminRouter.POST("/materialgroup", materialGroupController.Post)
		adminRouter.PUT("/materialgroup", materialGroupController.Put)
		adminRouter.DELETE("/materialgroup/:id", materialGroupController.Delete)

		adminRouter.GET("/user", userController.GetAll)
		adminRouter.POST("/user", userController.Post)
		adminRouter.PUT("/user", userController.Put)
		adminRouter.DELETE("/user", userController.Delete)
		adminRouter.PUT("/user/center", userController.Center)
		adminRouter.POST("/user/updatePass/", userController.Pass)
		adminRouter.POST("/user/updateAvatar", userController.Avatar)

		adminRouter.GET("/dept", deptController.GetAll)
		adminRouter.POST("/dept", deptController.Post)
		adminRouter.PUT("/dept", deptController.Put)
		adminRouter.DELETE("/dept", deptController.Delete)

		adminRouter.GET("/dict", dictController.GetAll)
		adminRouter.POST("/dict", dictController.Post)
		adminRouter.PUT("/dict", dictController.Put)
		adminRouter.DELETE("/dict/:id", dictController.Delete)

		adminRouter.GET("/dictDetail", dictDetailController.GetAll)
		adminRouter.POST("/dictDetail", dictDetailController.Post)
		adminRouter.PUT("/dictDetail", dictDetailController.Put)
		adminRouter.DELETE("/dictDetail/:id", dictDetailController.Delete)

		adminRouter.GET("/roles/:id", roleController.GetOne)
		adminRouter.GET("/roles", roleController.GetAll)
		adminRouter.POST("/roles", roleController.Post)
		adminRouter.PUT("/roles", roleController.Put)
		adminRouter.DELETE("/roles", roleController.Delete)
		adminRouter.PUT("/roles/menu", roleController.Menu)

		adminRouter.GET("/logs", logController.GetAll)
		adminRouter.DELETE("/logs", logController.Delete)

		adminRouter.GET("/menu/build", menuController.Build)
		adminRouter.GET("/menu/listtree", menuController.GetTree)
		adminRouter.GET("/menu", menuController.GetAll)
		adminRouter.POST("/menu", menuController.Post)
		adminRouter.PUT("/menu", menuController.Put)
		adminRouter.DELETE("/menu", menuController.Delete)
	}

	picController := admin.StorePicController{}
	shopRouter := r.Group("/shop")
	shopRouter.Use(middleware.Jwt()).Use(middleware.Log())
	{
		shopRouter.GET("/product", picController.GetAll)
		shopRouter.GET("/product/info/:id", picController.GetInfo)
		shopRouter.POST("/product/addOrSave", picController.Post)
		shopRouter.DELETE("/product/:id", picController.Delete)

	}

	return r
}
