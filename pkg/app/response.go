package app

import (
	"github.com/KenmyZhang/pic_site_admin_server/pkg/constant"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type BaseResponse struct {
	Code int    `json:"status"` // 状态码
	Msg  string `json:"msg"`    // 提示信息
}

type Response struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	switch errCode.(type) {
	case int:
		msg := ""
		if data != nil {
			if val, ok := data.(string); ok {
				msg = "," + val
			}
		}
		intCode := errCode.(int)
		g.C.JSON(httpCode, Response{
			Code: intCode,
			Msg:  constant.GetMsg(intCode) + msg,
			Data: data,
		})

	case string:
		strCode := errCode.(string)
		g.C.JSON(httpCode, Response{
			Code: 9999,
			Msg:  strCode,
			Data: data,
		})
	}

	return
}

// Response setting gin.JSON
func (g *Gin) ResponsePage(httpCode int, errCode interface{}, data interface{}, total, totalPage int) {
	intCode := errCode.(int)
	g.C.JSON(httpCode, ResponsePage{
		Code:      intCode,
		Msg:       constant.GetMsg(intCode),
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
	return
}
