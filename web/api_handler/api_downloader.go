package api_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Download_UpdateFile(c *fiber.Ctx) error {
	// /file/:winver
	return c.Download(fmt.Sprintf("./file/%s.msu", c.Params("winver")))
}
