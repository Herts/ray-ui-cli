package controllers

import "os/exec"

func RestartV2ray() {
	exec.Command("systemctl", "restart", "v2ray")
}

func NginxReload() {
	exec.Command("nginx", "-s reload")
}

func QueryStats() {
	exec.Command("v2ctl",)
}