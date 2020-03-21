package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os/exec"
	"strings"
)

type SystemController struct {
	beego.Controller
}

type DataConsumed struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
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
	cmd := exec.Command(_v2ctl, "api", "--server", "127.0.0.1:8144", "StatsService.QueryStats", "pattern: \"\"")
	log.Println(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err, string(output))
		return
	}

	stats := strings.ReplaceAll(string(output), "\n>", "\n},")
	log.Println(stats)
	stats = strings.ReplaceAll(stats, "stat: <", "{")
	stats = strings.ReplaceAll(stats, "value", ",\"value\"")
	stats = strings.ReplaceAll(stats, "name", "\"name\"")
	stats = fmt.Sprint("[", stats[:len(stats)-3], "]")
	log.Println(stats)

	err = json.Unmarshal([]byte(stats), &data)
	if err != nil {
		log.Println(err)
	}
	return data, nil
}
