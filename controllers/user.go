package controllers

import (
	"github.com/Herts/ray-ui-cli/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *UserController) GetUserDataPage() {
	data := models.GetAllDataConsumed()
	// TODO: the form to display data should be changed
	for _, d := range data {
		d.UpDataConsumed /= 10e6
		d.DownDataConsumed /= 10e6
	}
	c.Data["data"] = data
	c.Layout = "layout.html"
	c.TplName = "userdata.html"
}

func (c *UserController) GetUserPage() {
	//models.UpdateDataConsumedInUser()
	users := models.GetAllUser()
	c.Data["data"] = users
	c.Layout = "layout.html"
	c.TplName = "user.html"
}
