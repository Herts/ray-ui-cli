package controllers

import (
	"github.com/astaxie/beego"
	"os/exec"
)

type SystemController struct {
	beego.Controller
}

func (c *SystemController) RestartV2ray() {
	output, err := exec.Command("systemctl", "restart", "v2ray").Output()
	if err != nil {
		c.Data["json"] = response{
			Message: err.Error(),
		}
		return
	}
	c.Data["json"] = response{
		Message: string(output),
	}
}

func (c *SystemController) NginxReload() {
	exec.Command("nginx", "-s reload")
}

func (c *SystemController) QueryStats() {
	exec.Command("v2ctl")
}
