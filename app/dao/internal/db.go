package internal

import (
	"beegoframe/app/model"
	"beegoframe/config"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DB orm.Ormer

func init() {
	if DB != nil {
		return
	}
	address := config.Params.Database.Address
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", address)
	orm.RegisterModel(
		new(model.Account),
		new(model.Reader),
	)
	if config.Params.Database.Sync {
		orm.RunSyncdb("default", false, true)
	}
	orm.Debug = config.Params.Database.Debug
	DB = orm.NewOrm()
}
