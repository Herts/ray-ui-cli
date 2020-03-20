package models

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
	Port           int            `json:"port"`
	Protocol       string         `json:"protocol"`
	Settings       interface{}    `json:"settings"`
	StreamSettings StreamSettings `json:"streamSettings"`
}

type VmessClient struct {
	ID      string `json:"id"`
	Level   int    `json:"level"`
	AlterID int    `json:"alterId"`
	Email   string `json:"email"`
}

type VmessInboundConfiguration struct {
	Clients                   []VmessClient `json:"clients"`
	Default                   VmessDefault  `json:"default"`
	Detour                    VmessDetour   `json:"detour"`
	DisableInsecureEncryption bool          `json:"disableInsecureEncryption"`
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
	IP          []string `json:"ip"`
	OutboundTag string   `json:"outboundTag"`
}

type Routing struct {
	Rules []Rules `json:"rules"`
}
