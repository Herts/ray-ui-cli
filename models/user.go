package models

import (
	"github.com/Herts/ray-ui-cli/myutils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model   `json:"-"`
	Email        string `gorm:"primary_key" json:"email"`
	UserId       string `json:"userId"`
	Enabled      bool   `json:"enabled"`
	DataConsumed int64  `json:"dataConsumed"`
	Level        int    `json:"level,string"`
	AlterID      int    `json:"alterId,string"`
}

type RemoteServer struct {
	gorm.Model `json:"-"`
	NickName   string
	Address4   string
	ServerName string `json:"serverName"`
	Mask       string `json:"mask"`
	Host       string
	TLSName    string
	Port       int `json:"port,string"`
	Provider   string
	Price      float64
	RayPort    int    `json:"rayPort,string"`
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
	loc, _ := time.LoadLocation("Asia/Shanghai")

	for _, stat := range stats {
		if strings.HasPrefix(stat.Name, "user") {
			info := strings.Split(stat.Name, ">>>")
			email := info[1]
			emails = append(emails, email)
			ud := GetUserDataOneDay(email, time.Now().In(loc).Format("2006-01-02"))
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

func GetAllDataConsumedInRange(startDate, endDate *time.Time) (uds []*UserData) {
	var query *gorm.DB
	if startDate != nil {
		start := startDate.Format("2006-01-02")
		query = db.Where("DATE(date) >= ?", start)
	}
	if endDate != nil {
		end := endDate.Format("2006-01-02")
		query = db.Where("DATE(date) <= ?", end)

	}
	query.Find(&uds)
	return
}

func UpdateDataConsumedInUser() {
	users := GetAllUser()
	type result struct {
		Up   int64
		Down int64
	}
	var r result
	for _, u := range users {
		db.Model(&UserData{}).Select("sum(up_data_consumed) as up, sum(down_data_consumed) as down").Scan(&r)
		u.DataConsumed = r.Up + r.Down
		db.Save(u)
	}
}
