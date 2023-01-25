package api_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Download_Chartjs(c *fiber.Ctx) error {
	// http://localhost/cdn/chartjs
	return c.Download("./js/Chart.bundle.min.js")
}

func Download_UpdateDownloader(c *fiber.Ctx) error {
	return c.Download("./files/win_update.zip")
}

func Download_updatefile(c *fiber.Ctx) error {
	// /winver/:winver
	return c.Download(fmt.Sprintf("./files/%s.msu", c.Params("winver")))
}
