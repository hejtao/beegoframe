package controller

import (
	"beegoframe/app/shared"
	"beegoframe/app/system/index/internal/service"
)

var Reader = new(readerController)

type readerController struct {
	shared.Controller
}

func (c *readerController) Post() {
	if err := service.Reader.Create(); err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", nil)
	}
}

func (c *readerController) Get() {
	list, err := service.Reader.Get()
	if err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", list)
	}
}

func (c *readerController) Put() {
	if err := service.Reader.Update(); err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", nil)
	}
}

func (c *readerController) Delete() {
	if err := service.Reader.Delete(); err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", nil)
	}
}
