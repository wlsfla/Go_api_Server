package models

import "github.com/gofiber/fiber/v2"

type Hostinfo struct {
	Host_ip   string
	Host_name string
	Winver    string
	Build_ver string
	Result    string
}

func New_Hostinfo(c *fiber.Ctx) *Hostinfo {
	return &Hostinfo{
		c.IP(),
		c.Query("host_name"),
		c.Query("winver"),
		c.Query("buildver"),
		c.Query("result"),
	}
}
