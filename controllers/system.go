package controllers

import (
	"github.com/astaxie/beego"
)

type SystemController struct {
	beego.Controller
}

func (c *SystemController) GetSystemPage() {
	c.TplName = "system.html"
	c.Layout = "layout.html"
}

func (c *SystemController) GenerateApiToken() {
	//	TODO: implement this
}
