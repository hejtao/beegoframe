package internal

import "github.com/beego/beego/v2/client/orm"

func NewRow() orm.Params {
	var row orm.Params
	return row
}

func NewRows() []orm.Params {
	var rows []orm.Params
	return rows
}

func NewColumn() orm.ParamsList {
	var column orm.ParamsList
	return column
}
