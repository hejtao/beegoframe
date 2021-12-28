package internal

import (
	"beegoframe/config"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/olivere/elastic/v7"
	"time"
)

const interval = 10 * time.Second

var client *elastic.Client

var prefix string

func init() {
	fmt.Println("Connecting to the ES cluster...")

	var err error
	address := config.Params.Es.Address
	username := config.Params.Es.Username
	password := config.Params.Es.Password
	prefix = config.Params.Es.Prefix
	if client, err = elastic.NewClient(
		elastic.SetURL(address),
		elastic.SetBasicAuth(username, password),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(true),
		elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(interval),
	); err != nil {
		logs.Error(err)
		return
	}

	fmt.Println("Connect to the ES cluster successfully!")
	fmt.Println()
}
