package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"log"
	"os"
)

type RayConfig struct {
	Stats     interface{} `json:"stats"`
	API       API         `json:"api"`
	Policy    Policy      `json:"policy"`
	Log       Log         `json:"log"`
	Inbounds  []Inbounds  `json:"inbounds"`
	Outbounds []Outbounds `json:"outbounds"`
	Routing   Routing     `json:"routing"`
}

type API struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

type Stats struct {
	StatsUserUplink   bool `json:"statsUserUplink"`
	StatsUserDownlink bool `json:"statsUserDownlink"`
}

type System struct {
	StatsInboundUplink   bool `json:"statsInboundUplink"`
	StatsInboundDownlink bool `json:"statsInboundDownlink"`
}

type Policy struct {
	Levels map[string]Stats `json:"levels"`
	System System           `json:"system"`
}
type Log struct {
	Loglevel string `json:"loglevel"`
	Access   string `json:"access"`
	Error    string `json:"error"`
}

type Inbounds struct {
	Tag            string          `json:"tag"`
	Port           int             `json:"port"`
	Protocol       string          `json:"protocol"`
	Settings       interface{}     `json:"settings"`
	StreamSettings *StreamSettings `json:"streamSettings,omitempty"`
}

type VmessClient struct {
	ID      string `json:"id"`
	Level   int    `json:"level"`
	AlterID int    `json:"alterId"`
	Email   string `json:"email"`
}

type VmessInboundConfiguration struct {
	Clients                   []VmessClient `json:"clients"`
	Default                   *VmessDefault `json:"default,omitempty"`
	Detour                    *VmessDetour  `json:"detour,omitempty"`
	DisableInsecureEncryption bool          `json:"disableInsecureEncryption,omitempty"`
}

type VmessDefault struct {
	Level   int `json:"level"`
	AlterID int `json:"alterId"`
}

type VmessDetour struct {
	To string `json:"to"`
}

type StreamSettings struct {
	Network string `json:"network"`
}

type Outbounds struct {
	Protocol string      `json:"protocol"`
	Settings interface{} `json:"settings"`
	Tag      string      `json:"tag,omitempty"`
}

type Rules struct {
	Type        string   `json:"type"`
	IP          []string `json:"ip,omitempty"`
	OutboundTag string   `json:"outboundTag"`
}

type Routing struct {
	Rules []Rules `json:"rules"`
}

func ReloadConfig() *RayConfig {
	configFile := beego.AppConfig.String("v2rayconfig")
	//configFile := "tests/config.json"
	if len(configFile) == 0 {
		log.Println(errors.New("config not found"))
	}
	f, err := os.Open(configFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	var config RayConfig
	err = json.NewDecoder(f).Decode(&config)
	log.Printf("now config -> %v Err: %v\n", config, err)
	var users []*User
	db.Where(User{Enabled: true}).Find(&users)

	var clients []VmessClient
	for _, user := range users {
		client := VmessClient{
			ID:      user.UserId,
			Level:   user.Level,
			AlterID: user.AlterID,
			Email:   user.Email,
		}
		clients = append(clients, client)
	}

	vmess := VmessInboundConfiguration{
		Clients: clients,
	}

	config.Inbounds[0].Settings = vmess

	jsonConfig, err := json.Marshal(config)
	log.Println(string(jsonConfig))
	return &config

}
