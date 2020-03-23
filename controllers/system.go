package controllers

import (
	"../models"
	"../myutils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os/exec"
	"text/template"
)

type SystemController struct {
	beego.Controller
}

func (c *SystemController) RestartV2ray() {
	models.UpdateDataConsumed()
	c.ExecuteCmd("sudo", "systemctl", "restart", "v2ray")
}

func (c *SystemController) ExecuteCmd(command ...string) {
	cmd := exec.Command(command[0], command[1:]...)
	logs.Info(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.Data["json"] = response{
			Message: fmt.Sprint(string(output), err.Error()),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = response{
		Message: string(output),
	}
	c.ServeJSON()
}

func (c *SystemController) NginxReload() {
	c.ExecuteCmd("nginx", "-s", "reload")
}

func (c *SystemController) QueryStats() {
	exec.Command("v2ctl")
}

func (c *SystemController) ReGenConfig() {
	config, err := models.ReloadConfig()
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = response{Data: config}
	c.ServeJSON()
}

func (c *SystemController) CertbotGetCert() {
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

func (c *SystemController) GetRawStats() {
	stats, err := myutils.GetStatistics(false)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = response{Data: stats}
	c.ServeJSON()
}

func (c *SystemController) GenNginxConfig() {
	tpl, err := template.ParseFiles("conf/nginx.tpl")
	if err != nil {
		logs.Error(err)
		return
	}
	var server models.RemoteServer
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &server)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = response{Message: err.Error()}
		return
	}
	server.Port = 80

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, server)
	if err != nil {
		logs.Error(err)
	}
	config := buffer.Bytes()
	nginxDir := beego.AppConfig.DefaultString("nginxdir", "/etc/nginx/site-enabled/")
	err = ioutil.WriteFile(nginxDir+server.ServerName, buffer.Bytes(), 0644)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = response{Message: err.Error()}
		return
	}

	c.Data["json"] = response{Data: string(config)}
	c.ServeJSON()
}
