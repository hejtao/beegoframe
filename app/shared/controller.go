package shared

import (
	"beegoframe/pkg/encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"reflect"
)

var errInvalidContainerType = errors.New("invalid container type, need ptr struct")

type Controller struct {
	web.Controller
}

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *Controller) SetResp(hCode int, message string, data interface{}) {
	c.Ctx.Output.SetStatus(hCode)
	c.Data["json"] = Resp{message, data}
	c.ServeJSON()
}

func (c *Controller) ParseQuery(container interface{}) error {
	t := reflect.TypeOf(container)
	if t.Kind() != reflect.Ptr {
		return errInvalidContainerType
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return errInvalidContainerType
	}
	v := reflect.ValueOf(container).Elem()
	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		fieldPtr := v.Field(i).Addr().Interface()
		if err := c.Ctx.Input.Bind(fieldPtr, jsonTag); err != nil {
			return err
		}
	}
	err := valid(container)
	return err
}

func (c *Controller) ParseBody(container interface{}) error {
	t := reflect.TypeOf(container)
	if t.Kind() != reflect.Ptr {
		return errInvalidContainerType
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, container); err != nil {
		return err
	}
	err := valid(container)
	return err
}

func valid(md interface{}) (err error) {
	v := validation.Validation{}
	b, err := v.Valid(md)
	if err != nil {
		return err
	}
	errMsg := ""
	if !b {
		for _, err := range v.Errors {
			errMsg += fmt.Sprintf(", %s:%s", err.Field, err.Message)
		}
		err = errors.New(errMsg[1:])
	}
	return
}
