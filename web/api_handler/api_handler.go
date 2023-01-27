package api_handler

import (
	"app/DBConfig"
	"app/models"

	"github.com/gofiber/fiber/v2"
)

var (
	Winverlist map[string]string
)

func Init() {
	initWinverList()
}

func getBuildVer(c *fiber.Ctx) error {
	result := fiber.Map{
		"status":   0,
		"buildver": 0,
	}

	resp := c.AllParams()
	if resp == nil {
		return c.JSON(result)
	}

	buildver, exists := Winverlist[resp["winver"]]

	if !exists {
		return c.JSON(result)
	}

	result["status"] = 1
	result["buildver"] = buildver

	return c.JSON(result)
}

func initWinverList() {
	Winverlist = make(map[string]string)

	list := []models.Target_winvers{}
	DBConfig.DBconn.Find(&list)

	for _, v := range list {
		Winverlist[v.Winver] = v.Buildver
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

func Apiv2Router() *fiber.App {
	app := fiber.New()

	app.Post("/insert/updatelog", insertinfo)
	app.Get("/winver/:winver", getBuildVer)

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
// POST http://localhost/api/v2/updatelog
// content-type: application/json

// {
//     "Host_ip": "1.1.1.1",
//     "Host_name": "test_name",
//     "Winver": "22H1",
//     "Buildver": "19044.2486",
//     "Result": 0
// }
