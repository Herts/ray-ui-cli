package controllers

import (
	"github.com/astaxie/beego"
	"os/exec"
)

type SystemController struct {
	beego.Controller
}

func (c *SystemController) RestartV2ray() (out string, err error) {
	output, err := exec.Command("systemctl", "restart", "v2ray").Output()
	if err != nil {
		return
	}
	out = string(output)
	return
}

func (c *SystemController) NginxReload() {
	exec.Command("nginx", "-s reload")
}

func (c *SystemController) QueryStats() {
	exec.Command("v2ctl")
}
