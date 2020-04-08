// @APIVersion 1.0.0
// @Title ray-ui-cli API
// @Security ApiKeyAuth
// @SecurityDefinition ApiKeyAuth apiKey Authorization header "Authorization token set in conf/admin.conf"
package routers

import (
	"github.com/Herts/ray-ui-cli/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserApiController{}),
		),
		beego.NSNamespace("/system",
			beego.NSInclude(&controllers.SystemApiController{}),
		),
	)
	beego.AddNamespace(ns)

	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.InsertFilter("/*", beego.BeforeRouter, InitFilter())
	beego.Router("/", &controllers.MainController{})
	beego.Router("/html/login", &controllers.MainController{}, "get:LoginPage")
	beego.Router("/html/login", &controllers.MainController{}, "post:Login")

	beego.Router("/html/userdata", &controllers.UserController{}, "get:GetUserDataPage")
	beego.Router("/html/users", &controllers.UserController{}, "get:GetUserPage")
	beego.Router("/html/system", &controllers.SystemController{}, "get:GetSystemPage")
}

func InitFilter() beego.FilterFunc {
	auth := beego.AppConfig.String("Authorization")
	return func(ctx *context.Context) {
		if strings.HasPrefix(ctx.Input.URL(), "/html/login") {
			return
		}
		if len(auth) != 0 {
			if ctx.Input.Header("Authorization") == auth {
				return
			}
		}
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/html/login")
		}
	}
}
