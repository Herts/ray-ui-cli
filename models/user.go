package models

import (
	"../myutils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model   `json:"-"`
	Email        string  `gorm:"primary_key" json:"email"`
	UserId       string  `json:"userId"`
	Enabled      bool    `json:"enabled"`
	DataConsumed float64 `json:"dataConsumed"`
	Level        int     `json:"level"`
	AlterID      int     `json:"alterId"`
}

type RemoteServer struct {
	gorm.Model `json:"-"`
	NickName   string
	Address4   string
	ServerName string `json:"serverName"`
	Mask       string `json:"mask"`
	Host       string
	TLSName    string
	Port       int `json:"port"`
	Provider   string
	Price      float64
	RayPort    int    `json:"rayPort"`
	Region     string `gorm:"primary_key"`
	Index      int    `gorm:"primary_key;auto_increment:false"`
	ExpiresOn  time.Time
}

type UserServer struct {
	gorm.Model `json:"-"`
	NickName   string `gorm:"primary_key"`
	Region     string `gorm:"primary_key"`
	Index      int    `gorm:"primary_key;auto_increment:false"`
}

type UserData struct {
	gorm.Model       `json:"-"`
	Email            string `gorm:"primary_key" json:"email"`
	Date             string `gorm:"primary_key" json:"date"`
	UpDataConsumed   int64  `json:"upDataConsumed"`
	DownDataConsumed int64  `json:"downDataConsumed"`
}

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "local.db")
	if err != nil {
		logs.Error(err)
	}
	dbDebug, err := beego.AppConfig.Bool("dbdebug")
	if err != nil {
		dbDebug = false
	}
	db.LogMode(dbDebug)
	dbInit, err := beego.AppConfig.Bool("dbinit")
	if err != nil {
		dbInit = false
	}
	if !dbInit {
		db.AutoMigrate(&User{})
		db.AutoMigrate(&UserData{})
	}
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

func GetUserDataOneDay(email string, day string) *UserData {
	var ud UserData
	db.FirstOrInit(&ud, UserData{Email: email, Date: day})
	return &ud
}

func SaveUserData(ud *UserData) {
	db.Save(ud)
}

func UpdateDataConsumed() (emails []string) {
	stats, err := myutils.GetStatistics(true)
	if err != nil {
		logs.Error(err)
		return
	}
	for _, stat := range stats {
		if strings.HasPrefix(stat.Name, "user") {
			info := strings.Split(stat.Name, ">>>")
			email := info[1]
			emails = append(emails, email)
			ud := GetUserDataOneDay(email, time.Now().Format("2006-01-02"))
			if strings.HasSuffix(stat.Name, "uplink") {
				ud.UpDataConsumed += stat.Value
			} else {
				ud.DownDataConsumed += stat.Value
			}
			SaveUserData(ud)
		}
	}
	return
}

func GetAllDataConsumed() (uds []*UserData) {
	db.Find(&uds)
	return
}
