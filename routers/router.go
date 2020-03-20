package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/add", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/api/user/update", &controllers.UserController{}, "put:UpdateUser")
	beego.Router("/api/user/list", &controllers.UserController{}, "get:ListAllUsers")
}
