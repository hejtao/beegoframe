package index

import (
	"beegoframe/app/system/index/internal/controller"
	"github.com/beego/beego/v2/server/web"
)

func initRouter() {
	web.Router("/api/account", controller.Account)
	web.Router("/api/reader", controller.Reader)
}
