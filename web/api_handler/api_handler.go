package api_handler

import (
	"app/DBConfig"
	"app/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	Winverlist map[string]string
)

func Init() {
	initWinverList()
	// init_winverlist()
}

func GetBuildVer(c *fiber.Ctx) error {
	resp := c.AllParams()
	if resp == nil {
		return c.Status(400).JSON(fiber.Map{
			"result":   0,
			"buildver": nil,
		})
	}

	fmt.Println(Winverlist)
	buildver := Winverlist[resp["winver"]]

	return c.JSON(fiber.Map{
		"result":   1,
		"buildver": buildver,
	})
}

func initWinverList() {
	// Winverlist :=

	list := []models.Target_winvers{}
	DBConfig.DBconn.Find(&list)
}

func insertinfo(c *fiber.Ctx) error {
	updatelog := new(models.Update_Logs)
	if err := c.BodyParser(updatelog); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	DBConfig.DBconn.Create(&updatelog)

	return c.Status(200).JSON(updatelog)
}

func Apiv2() *fiber.App {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("api test")
	})

	app.Post("/updatelog", insertinfo)
	app.Get("/winver/:winver", GetBuildVer)

	return app
}

// app.Get("/insert/info", api_handler.InsertHostinfo)
// app.Get("/update/info", api_handler.UpdateHostinfo)

// /updateinfo?
// 		host_name=${host_name}&
//		winver=${winver}&
//		buildver=${buildver}&
//		created_time=${created_time}&
//		updated_time=${updated_time}&
//		result=${result}

// test
// POST http://localhost/api/v2/updatelog HTTP/1.1
// content-type: application/json

// {
//     "Host_ip": "1.1.1.1",
//     "Host_name": "test_name",
//     "Winver": "22H1",
//     "Buildver": "19044.2486",
//     "Result": 0
// }
