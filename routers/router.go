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

	beego.Router("/api/system/restartV2ray", &controllers.SystemController{}, "get:RestartV2ray")
	beego.Router("/api/system/nginxReload", &controllers.SystemController{}, "get:NginxReload")
	beego.Router("/api/system/restartV2ray", &controllers.SystemController{}, "get:RestartV2ray")

}
