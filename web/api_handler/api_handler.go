package api_handler

import (
	"app/DBConfig"
	"app/common"
	"app/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	Winverlist        map[string]string
	updatefileUrlList map[string]string
)

func Init() {
	initWinverList()
}

// Request OS Version
// Return Matched Build Version
func getBuildVer(c *fiber.Ctx) error {
	result := models.Winver_info{
		Status:   0,
		Winver:   "",
		Buildver: "",
		Url:      "",
	}

	resp := c.AllParams()
	if resp == nil {
		return c.JSON(result)
	}

	result.Winver = resp["winver"]
	buildver, exists := Winverlist[result.Winver]

	if !exists {
		return c.JSON(result)
	}

	result.Status = 1
	result.Buildver = buildver
	result.Url = updatefileUrlList[result.Winver]

	return c.JSON(result)
}

func initWinverList() {
	Winverlist = make(map[string]string)
	updatefileUrlList = make(map[string]string)

	list := []models.Target_winvers{}
	DBConfig.DBconn.Find(&list)

	for _, v := range list {
		Winverlist[v.Winver] = v.Buildver
		updatefileUrlList[v.Winver] = fmt.Sprintf("http://" + common.Server_ip + "/static/files/" + v.Winver + ".msu")
	}
}

func insertinfo(c *fiber.Ctx) error {
	updatelog := new(models.Update_Logs)
	if err := c.BodyParser(updatelog); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	DBConfig.DBconn.Create(&updatelog)

	return c.Status(200).JSON(updatelog)
}

// API v2 Main Routing
func Apiv2Router() *fiber.App {
	app := fiber.New()

	app.Post("/insert/updatelog", insertinfo)
	app.Get("/winver/:winver", getBuildVer)

	return app
}
