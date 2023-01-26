package models

import (
	"time"

	"gorm.io/gorm"
)

type Update_Log struct {
	gorm.Model

	Host_ip      string    `json:"host_ip"`
	Host_name    string    `json:"host_name"`
	Winver       string    `json:"winver"`
	Build_ver    string    `json:"buildver"`
	Updated_time time.Time `json:"updated_time"`
	Result       int       `json:"title"`
}

type Target_winver struct {
	gorm.Model

	Winver   string `json:"winver"`
	Buildver string `json:"buildver"`
	KbNumber string `json:"kbNumber"`
}
