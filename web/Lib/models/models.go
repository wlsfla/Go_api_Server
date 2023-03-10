package models

import (
	"time"
)

type Update_Logs struct {
	Host_ip      string    `json:"host_ip"`
	Host_name    string    `json:"host_name"`
	Winver       string    `json:"winver"`
	Buildver     string    `json:"buildver"`
	Updated_time time.Time `json:"updated_time" gorm:"-:all;autoCreateTime"`
	Result       int       `json:"result"`
}

type Target_winvers struct {
	Winver   string `json:"winver"`
	Buildver string `json:"buildver"`
	KbNumber string `json:"kbNumber"`
}

type Winver_info struct {
	Status   int    `json:"status"`
	Winver   string `json:"winver"`
	Buildver string `json:"buildver"`
	Url      string `json:"url"`
}
