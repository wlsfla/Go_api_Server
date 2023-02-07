package ConnManager

import (
	"app/Lib/api_router/Api_handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	MaxConnCount  uint
	CurrConnCount uint
)

func init() {
	fmt.Println("\t[*] init ConnManager")

	MaxConnCount = 3
}

func SetMaxConnCount(value uint) {
	if value > 0 {
		MaxConnCount = value
		Showinfo()
	}
}

func Showinfo() {
	fmt.Printf("[***] CurrConnCount: %d | MaxConnCount: %d\n", CurrConnCount, MaxConnCount)
}

func GetConnInfo(c *fiber.Ctx) error {
	// /connpool/info

	return c.JSON(fiber.Map{
		"status":    1,
		"CurrCount": CurrConnCount,
		"MaxCount":  MaxConnCount,
	})
}

func isCanDownload() bool {
	if MaxConnCount >= CurrConnCount {
		return false
	} else {
		return true
	}
}

func Getupdatefile(c *fiber.Ctx) error {
	// /file/:winver
	if isCanDownload() {
		winver := c.Params("winver")

		// fmt.Sprintf("http://" + common.Server_ip + "/static/files/" + v.Winver + ".msu")

		return c.JSON(fiber.Map{
			"status": 1,
			"URL":    Api_handler.GetFileUrl(winver),
		})
	}

	return c.JSON(fiber.Map{
		"status": 0,
		"URL":    "",
	})
}

func ChangeConnPool(c *fiber.Ctx) error {
	// /connpool/max/:value

	v, _ := c.ParamsInt("value")

	SetMaxConnCount(uint(v))

	return c.JSON(fiber.Map{
		"status":    1,
		"CurrCount": CurrConnCount,
		"MaxCount":  MaxConnCount,
	})
}
