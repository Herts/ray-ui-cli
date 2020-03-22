package controllers

import (
	"../models"
	"../myutils"
	"github.com/astaxie/beego"
	"os/exec"
)

type SystemController struct {
	beego.Controller
}

func (c *SystemController) RestartV2ray() {
	c.ExecuteCmd("sudo", "systemctl", "restart", "v2ray")
}

func (c *SystemController) ExecuteCmd(command ...string) {
	output, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		c.Data["json"] = response{
			Message: err.Error(),
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
	email := "lhz007563@gmail.com"
	domainName := "test.222422.xyz"
	c.ExecuteCmd("sudo", "certbot", "--nginx", "-m", email, "--agree-tos", "--no-eff-email", "-d",
		domainName, "--no-redirect")
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
