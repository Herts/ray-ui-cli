package main

import (
	"./models"
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/jasonlvhit/gocron"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

}

func timelyTask() {
	interval := beego.AppConfig.DefaultInt64("dataupdateinterval", 10)

	gocron.Every(uint64(interval)).Minutes().From(gocron.NextTick()).Do(models.UpdateDataConsumed)
	gocron.Start()
}

func main() {
	timelyTask()
	models.InitDB()
	beego.Run()
}
