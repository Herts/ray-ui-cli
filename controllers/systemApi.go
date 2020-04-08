package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Herts/ray-ui-cli/models"
	"github.com/Herts/ray-ui-cli/myutils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os/exec"
	"text/template"
)

type SystemApiController struct {
	beego.Controller
}

// @Success 200 {response} response
// @router /restartV2ray [get]
func (c *SystemApiController) RestartV2ray() {
	models.UpdateDataConsumed()
	c.ExecuteCmd("systemctl", "restart", "v2ray")
}

func (c *SystemApiController) ExecuteCmd(command ...string) {
	cmd := exec.Command(command[0], command[1:]...)
	logs.Debug(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.Data["json"] = response{
			Message: fmt.Sprint(string(output), err.Error()),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = response{
		Message: fmt.Sprintf("No error occurred. %s", output),
	}
	c.ServeJSON()
}

// @Success 200 {response} response
// @router /nginxReload [get]
func (c *SystemApiController) NginxReload() {
	c.ExecuteCmd("nginx", "-s", "reload")
}

// @Success 200 {response} response
// @router /regenV2rayConfig [get]
func (c *SystemApiController) ReGenConfig() {
	config, err := models.ReloadConfig()
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = response{Data: config}
	c.ServeJSON()
}

// @Param body body server true "server"
// @Success 200 {response} response
// @router /certbotGetCert [post]
func (c *SystemApiController) CertbotGetCert() {
	type server struct {
		Email      string `json:"email"`
		ServerName string `json:"serverName"`
	}
	var s server
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &s)
	if err != nil {
		logs.Error(err)
	}
	c.ExecuteCmd("certbot", "--nginx", "-m", s.Email, "--agree-tos", "--no-eff-email", "-d",
		s.ServerName, "--no-redirect")
}

// @Success 200 {response} response
// @router /getRawStats [get]
func (c *SystemApiController) GetRawStats() {
	stats, err := myutils.GetStatistics(false)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = response{Data: stats}
	c.ServeJSON()
}

// @Param body body models.RemoteServer true "server"
// @Success 200 {response} response
// @router /genNginxConfig [post]
func (c *SystemApiController) GenNginxConfig() {
	tpl, err := template.ParseFiles("conf/nginx.tpl")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	var server models.RemoteServer
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &server)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	server.Port = 80

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, server)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	config := buffer.Bytes()
	nginxDir := beego.AppConfig.DefaultString("nginxdir", "/etc/nginx/site-enabled/")
	err = ioutil.WriteFile(nginxDir+server.ServerName, buffer.Bytes(), 0644)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = response{Data: string(config)}
	c.ServeJSON()
}
