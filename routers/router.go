package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.InsertFilter("/*", beego.BeforeRouter, InitFilter())
	beego.Router("/", &controllers.MainController{})
	beego.Router("/html/login", &controllers.MainController{}, "get:LoginPage")
	beego.Router("/html/login", &controllers.MainController{}, "post:Login")

	beego.Router("/api/user/add", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/api/user/update", &controllers.UserController{}, "put:UpdateUser")
	beego.Router("/api/user/list", &controllers.UserController{}, "get:ListAllUsers")
	beego.Router("/api/user/listData", &controllers.UserController{}, "get:ListAllDataConsumed")
	beego.Router("/api/user/updateData", &controllers.UserController{}, "get:UpdateDataConsumed")

	beego.Router("/html/userdata", &controllers.UserController{}, "get:GetUserDataPage")
	beego.Router("/html/system", &controllers.SystemController{}, "get:GetSystemPage")

	beego.Router("/api/system/restartV2ray", &controllers.SystemController{}, "get:RestartV2ray")
	beego.Router("/api/system/nginxReload", &controllers.SystemController{}, "get:NginxReload")
	beego.Router("/api/system/restartV2ray", &controllers.SystemController{}, "get:RestartV2ray")
	beego.Router("/api/system/regenV2ayConfig", &controllers.SystemController{}, "get:ReGenConfig")
	beego.Router("/api/system/getRawStats", &controllers.SystemController{}, "get:GetRawStats")
	beego.Router("/api/system/genNginxConfig", &controllers.SystemController{}, "post:GenNginxConfig")
	beego.Router("/api/system/certbotGetCert", &controllers.SystemController{}, "post:CertbotGetCert")
}

func InitFilter() beego.FilterFunc {
	return func(ctx *context.Context) {
		if strings.HasPrefix(ctx.Input.URL(), "/html/login") {
			return
		}

		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/html/login")
		}
	}
}
