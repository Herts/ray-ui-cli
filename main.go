package main

import (
	"encoding/json"
	"fmt"
	"github.com/Herts/ray-ui-cli/models"
	_ "github.com/Herts/ray-ui-cli/routers"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func init() {

}

func main() {
	models.InitDB()
	beego.Run()
}

func dbtest1() {
	db, err := gorm.Open("sqlite3", "local.db")
	if err != nil {
		log.Println(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&models.User{})

	user := models.User{
		Email:        "lhz007563@gmail.com",
		UserId:       uuid.New().String(),
		Enabled:      true,
		DataConsumed: 0,
	}
	db.Save(&user)
}

func test1() {
	var configObj models.RayConfig
	config, err := os.Open("tests/config.json")
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(config).Decode(&configObj)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(configObj)

	jsonObj, err := json.Marshal(configObj)
	fmt.Println(string(jsonObj))
}
