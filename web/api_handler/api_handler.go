package api_handler

import (
	"app/DBConfig"
	"app/models"

	"github.com/gofiber/fiber/v2"
)

var Winverlist map[string]string

func Init() {
	// winverlist = getTarget_winver()
	// init_winverlist()
}

func GetBuildVer(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"result":   1,
		"buildver": Winverlist[c.Params("winver")],
	})
}

func Test(c *fiber.Ctx) error {
	updatelog := new(models.Update_Log)
	updatelog.
	if err := c.BodyParser(updatelog); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	DBConfig.DBconn.Create(&updatelog)

	return c.Status(200).JSON(updatelog)
}
