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
	beego.Router("/api/user/listData", &controllers.UserController{}, "get:ListAllDataConsumed")
	beego.Router("/api/user/updateData", &controllers.UserController{}, "get:UpdateDataConsumed")
	beego.Router("/userdata", &controllers.UserController{}, "get:GetUserDataPage")

	beego.Router("/api/system/restartV2ray", &controllers.SystemController{}, "get:RestartV2ray")
	beego.Router("/api/system/nginxReload", &controllers.SystemController{}, "get:NginxReload")
	beego.Router("/api/system/restartV2ray", &controllers.SystemController{}, "get:RestartV2ray")
	beego.Router("/api/system/regenV2ayConfig", &controllers.SystemController{}, "get:ReGenConfig")
	beego.Router("/api/system/getRawStats", &controllers.SystemController{}, "get:GetRawStats")
	beego.Router("/api/system/genNginxConfig", &controllers.SystemController{}, "post:GenNginxConfig")
	beego.Router("/api/system/certbotGetCert", &controllers.SystemController{}, "post:CertbotGetCert")

}
