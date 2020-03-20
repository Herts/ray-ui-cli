package controllers

import (
	"../models"
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
	"os/exec"
	"strings"
)

type SystemController struct {
	beego.Controller
}

type DataConsumed struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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
	c.ExecuteCmd("nginx", "-s reload")
}

func (c *SystemController) QueryStats() {
	exec.Command("v2ctl")
}

func (c *SystemController) ReGenConfig() {
	config := models.ReloadConfig()
	c.Data["json"] = response{Data: config}
	c.ServeJSON()
}

func (c *SystemController) CertbotGetCert() {
	email := "lhz007563@gmail.com"
	domainName := "test.222422.xyz"
	c.ExecuteCmd("sudo", "certbot", "--nginx", "-m", email, "--agree-tos --no-eff-email", "-d",
		domainName, "--no-redirect")
}

func (c *SystemController) GetRawStats() {
	stats, err := GetStatistics()
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = response{Data: stats}
	c.ServeJSON()
}

func GetStatistics() (data []*DataConsumed, err error) {
	_v2ctl := "/usr/bin/v2ray/v2ctl"
	output, err := exec.Command(_v2ctl, "api", "--server=127.0.0.1:8144", "StatsService.QueryStats ''").Output()

	if err != nil {
		log.Println(err)
		return
	}

	stats := strings.ReplaceAll(string(output), "\n>", "\n}")
	stats = strings.ReplaceAll(stats, "stat: <", "{")
	stats = strings.ReplaceAll(stats, "value", ",\"value\"")
	stats = strings.ReplaceAll(stats, "name", "\"name\"")
	stats = fmt.Sprint("[", stats[:len(stats)-1], "]")

	err = json.Unmarshal([]byte(stats), data)
	if err != nil {
		log.Println(err)
	}
	return data, nil
}
