package ConnManager

import (
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
	}

	Showinfo()
}

func Showinfo() {
	fmt.Printf("[***] CurrConnCount: {%d} | MaxConnCount: {%d}", CurrConnCount, MaxConnCount)
}

func GetConnInfo(c *fiber.Ctx) error {
	// /connpool/info

	return c.JSON(fiber.Map{
		"status":    1,
		"CurrCount": CurrConnCount,
		"MaxCount":  MaxConnCount,
	})
}

func Getupdatefile(c *fiber.Ctx) error {

	return nil
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
