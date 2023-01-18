package api_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Download_Chartjs(c *fiber.Ctx) error {
	return c.Download("./js/Chart.bundle.min.js")
}

func Download_Downloader(c *fiber.Ctx) error {
	return c.Download("./file/win_update.zip")
}

func Download_updatefile(c *fiber.Ctx) error {
	// /winver/:winver
	return c.Download(fmt.Sprintf("./file/updatefile/%s.msu", c.Params("winver")))
}
