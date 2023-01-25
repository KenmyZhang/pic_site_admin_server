package dto

import "github.com/KenmyZhang/pic_site_admin_server/app/models/dto"

type UserQuery struct {
	dto.BasePage
	Sort    string
	Blurry  string
	Enabled bool
}
