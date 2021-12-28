package controller

import (
	"beegoframe/app/shared"
	"beegoframe/app/system/index/internal/service"
)

var Account = new(accountController)

type accountController struct {
	shared.Controller
}

func (c *accountController) Post() {
	if err := service.Account.Create(); err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", nil)
	}
}

func (c *accountController) Get() {
	list, err := service.Account.Get()
	if err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", list)
	}
}

func (c *accountController) Put() {
	if err := service.Account.Update(); err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", nil)
	}
}

func (c *accountController) Delete() {
	if err := service.Account.Delete(); err != nil {
		c.SetResp(shared.Hc500, err.Error(), nil)
	} else {
		c.SetResp(shared.Hc200, "", nil)
	}
}
