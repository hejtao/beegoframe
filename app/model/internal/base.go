package internal

import "time"

type Base struct {
	Id         int64     `json:"id"`
	Deleted    bool      `json:"deleted"`
	CreateTime time.Time `orm:"auto_now_add" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now" json:"update_time"`
}
