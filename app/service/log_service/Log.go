package log_service

import (
	"github.com/KenmyZhang/pic_site_admin_server/app/models"
	"github.com/KenmyZhang/pic_site_admin_server/app/models/vo"
)

type Log struct {
	Id int64

	PageNum  int
	PageSize int

	M *models.SysLog

	Ids []int64

	Des string
}

func (d *Log) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Des != "" {
		maps["description"] = d.Des
	}

	total, list := models.GetAllLog(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Log) Insert() error {
	return models.AddLog(d.M)
}

func (d *Log) Del() error {
	return models.DelBylog(d.Ids)
}
