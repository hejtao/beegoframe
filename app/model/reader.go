package model

import "beegoframe/app/model/internal"

type Reader struct {
	internal.Base
	Name      string `json:"name" orm:"size(16)"`
	Gender    string `json:"gender" orm:"size(8)"`
	Address   string `json:"address" orm:"size(64)"`
	Age       int8   `json:"age"`
	AccountId int64  `json:"account_id"`
}
