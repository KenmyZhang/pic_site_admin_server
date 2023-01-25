package listen

import (
	"fmt"

	"github.com/KenmyZhang/pic_site_admin_server/pkg/global"
)

func Setup() {
	var sub PSubscriber
	fmt.Printf(global.YSHOP_CONFIG.Redis.Host)
	conn := PConnect(global.YSHOP_CONFIG.Redis.Host, global.YSHOP_CONFIG.Redis.Password)
	sub.ReceiveKeySpace(conn)
	sub.Psubscribe()
}
