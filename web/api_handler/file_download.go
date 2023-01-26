package api_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Download_updatefile(c *fiber.Ctx) error {
	// /winver/:winver
	return c.SendString(
		fmt.Sprintf("./files/%s.msu", c.Params("winver")),
	)
}
