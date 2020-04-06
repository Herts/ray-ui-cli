package main

import (
	"github.com/Herts/ray-ui-cli/models"
	_ "github.com/Herts/ray-ui-cli/routers"
	"github.com/astaxie/beego"
	"github.com/jasonlvhit/gocron"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	timelyTask()
	models.InitDB()
}

func timelyTask() {
	interval := beego.AppConfig.DefaultInt64("dataupdateinterval", 10)
	gocron.Every(uint64(interval)).Minutes().From(gocron.NextTick()).Do(models.UpdateDataConsumed)
	gocron.Start()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
