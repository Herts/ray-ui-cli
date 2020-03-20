package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email        string  `gorm:"primary_key",json:"email"`
	UserId       string  `json:"userId"`
	Enabled      bool    `json:"enabled"`
	DataConsumed float64 `json:"dataConsumed"`
	Level        int     `json:"level"`
	AlterID      int     `json:"alterId"`
}

type RemoteServer struct {
	gorm.Model
	NickName   string
	Address4   string
	DomainName string
	Host       string
	TLSName    string
	Port       int
	Provider   string
	Price      float64
	Region     string `gorm:"primary_key"`
	Index      int    `gorm:"primary_key;auto_increment:false"`
	ExpiresOn  time.Time
}

type UserServer struct {
	gorm.Model
	NickName string `gorm:"primary_key"`
	Region   string `gorm:"primary_key"`
	Index    int    `gorm:"primary_key;auto_increment:false"`
}

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "local.db")
	if err != nil {
		log.Println(err)
	}
	db.LogMode(true)
	//db.AutoMigrate(&User{})
}

func AddUser(user *User) {
	user.Enabled = true
	if user.AlterID <= 0 {
		user.AlterID = 16
	}
	db.Save(user)
}

func GetUser(email string) *User {
	u := User{}
	db.Where(User{
		Email: email,
	}).First(&u)
	return &u
}

func UpdateUser(user *User) {
	if user.AlterID <= 0 {
		user.AlterID = 16
	}
	db.Save(user)
}

func GetAllUser() (users []*User) {
	db.Find(&users)
	return
}
