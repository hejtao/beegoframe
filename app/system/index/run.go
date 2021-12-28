package index

import (
	"beegoframe/config"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

func Run() {
	if config.Params.IndexServer.Enable {
		jwtKey = config.Params.IndexServer.JwtKey
		initFilter()
		initRouter()
		web.BConfig.CopyRequestBody = true
		port := config.Params.IndexServer.Port
		address := config.Params.IndexServer.Address
		web.Run(fmt.Sprintf("%s:%d", address, port))
	}
}
