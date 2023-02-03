package Api_handler

import (
	"app/Lib/DBConfig"
	"app/Lib/common"
	"app/Lib/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	Winverlist        map[string]string
	updatefileUrlList map[string]string
)

func init() {
	initWinverList()

}

func GetFileUrl(winver string) string {
	return updatefileUrlList[winver]
}

// Request OS Version
// Return Matched Build Version
func GetBuildVer(c *fiber.Ctx) error {
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
		updatefileUrlList[v.Winver] = fmt.Sprintf("http://" + common.Server_ip + "/api/v2/file/" + v.Winver)
	}
}

func Insertinfo(c *fiber.Ctx) error {
	updatelog := new(models.Update_Logs)
	if err := c.BodyParser(updatelog); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updatelog.Host_ip = c.IP()

	DBConfig.DBconn.Create(&updatelog)

	return c.Status(200).JSON(updatelog)
}
