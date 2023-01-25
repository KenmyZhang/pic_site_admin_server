package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KenmyZhang/pic_site_admin_server/app/service/pic_service"

	dto2 "github.com/KenmyZhang/pic_site_admin_server/app/service/pic_service/dto"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/app"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/constant"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/logging"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type StorePicController struct {
}

func (e *StorePicController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("is_show", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	categoryID, _ := strconv.ParseInt(c.Query("category_id"), 10, 64)
	picService := pic_service.Product{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
		CatId:    categoryID,
	}

	vo := picService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

func (e *StorePicController) GetInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	picService := pic_service.Product{
		Id: id,
	}
	vo := picService.GetProductInfo()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

func (e *StorePicController) Post(c *gin.Context) {
	var (
		dto  dto2.StorePic
		appG = app.Gin{C: c}
	)
	fmt.Println(fmt.Sprintf("dto:%+v", dto))
	httpCode, errCode, errMsg := app.BindAndValid(c, &dto)
	if errCode != constant.SUCCESS {
		logging.Error(errCode)
		appG.Response(httpCode, errCode, errMsg)
		return
	}
	picService := pic_service.Product{
		Dto: dto,
	}

	if err := picService.AddOrSaveProduct(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

func (e *StorePicController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)

	picService := pic_service.Product{Ids: ids}
	if err := picService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
