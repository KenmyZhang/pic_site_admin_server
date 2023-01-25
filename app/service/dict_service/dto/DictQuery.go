package dto

import "github.com/KenmyZhang/pic_site_admin_server/app/models/dto"

type DictQuery struct {
	dto.BasePage
	Blurry string
}
