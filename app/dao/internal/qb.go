package internal

import (
	"github.com/beego/beego/v2/client/orm"
)

type QueryBuilder interface {
	orm.QueryBuilder
}

func NewQueryBuilder() QueryBuilder {
	Builder, _ := orm.NewQueryBuilder("mysql")
	return Builder
}
