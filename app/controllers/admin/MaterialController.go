package admin

import (
	"net/http"

	"github.com/KenmyZhang/pic_site_admin_server/app/models"
	"github.com/KenmyZhang/pic_site_admin_server/app/service/material_service"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/app"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/constant"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/global"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/jwt"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/logging"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/upload"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"context"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 素材api
type MaterialController struct {
}

// @Title 素材列表
// @Description 岗位列表
// @Success 200 {object} app.Response
// @router / [get]
func (e *MaterialController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	groupId := com.StrTo(c.DefaultQuery("groupId", "-1")).MustInt64()
	name := c.DefaultQuery("blurry", "")
	materialService := material_service.Material{
		GroupId:  groupId,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := materialService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 素材添加
// @Description 素材添加
// @Success 200 {object} app.Response
// @router / [post]
func (e *MaterialController) Post(c *gin.Context) {
	var (
		model models.SysMaterial
		appG  = app.Gin{C: c}
	)
	httpCode, errCode, errMsg := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialService := material_service.Material{
		M: &model,
	}

	if err := materialService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 素材修改
// @Description 素材修改
// @Success 200 {object} app.Response
// @router / [put]
func (e *MaterialController) Put(c *gin.Context) {
	var (
		model models.SysMaterial
		appG  = app.Gin{C: c}
	)
	httpCode, errCode, errMsg := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialService := material_service.Material{
		M: &model,
	}

	if err := materialService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 素材删除
// @Description 素材删除
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *MaterialController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	//c.BindJSON(&ids)
	materialService := material_service.Material{Ids: ids}

	if err := materialService.Del(); err != nil {
		global.YSHOP_LOG.Error(err)
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 上传图像
// @Description 上传图像
// @Success 200 {object} app.Response
// @router /upload [post]
func (e *MaterialController) Upload(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, constant.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	//savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, constant.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, constant.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	// if err := TencentUploadFfile(src); err != nil {
	// 	logging.Warn(err)
	// 	appG.Response(http.StatusInternalServerError, constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
	// 	return
	// }

	imageUrl := upload.GetImageFullUrl(imageName)
	//imageSaveUrl := avePath + imageName
	//imageUrl := "https://pic.kimways.com/" + src

	appG.Response(http.StatusOK, constant.SUCCESS, imageUrl)

}

func TencentUploadFfile(filename string) error {
	u, _ := url.Parse("https://kimways-1253834929.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDwjpogMDbETaXsBaIG4g6Vkt00oQvfVYQ",
			SecretKey: "pX9e0BA0e7OTPAj9OSK4f8OUSErJafZR",
		},
	})
	// 通过文件流上传对象
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = c.Object.Put(context.Background(), filename, fd, nil)
	if err != nil {
		return err
	}
	return nil
}
