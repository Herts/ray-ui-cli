package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) LoginPage() {
	c.Layout = "layout.html"
	c.TplName = "login.html"
}

func (c *MainController) Login() {
	type user struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
		logs.Error(err)
		return
	}
	email := beego.AppConfig.String("email")
	pass := beego.AppConfig.String("password")
	if u.Email == email && u.Password == pass {
		c.SetSession("uid", 1)
		logs.Info("login success")
		c.Redirect("/html/userdata", 302)

	} else {
		c.Redirect("/html/login", 302)
	}
}
