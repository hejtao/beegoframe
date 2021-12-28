package model

import "beegoframe/app/model/internal"

type Account struct {
	internal.Base
	Email    string `orm:"size(32)"`
	Phone    string `orm:"size(16)"`
	Password string `orm:"size(32)"`
}
