package vo

import "github.com/KenmyZhang/pic_site_admin_server/app/models"

type LoginVo struct {
	Token string          `json:"token"`
	User  *models.SysUser `json:"user"`
}
