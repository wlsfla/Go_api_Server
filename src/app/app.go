package main

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()

	file, err := os.OpenFile("./log/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time}|${ip}|${port}|${status}|${method}|${path}\n",
		TimeFormat: "2006-01-02|15:04:05",
		TimeZone:   "Asia/Seoul",
		Output:     io.MultiWriter(file, os.Stdout),
	}))

	app.Get("/update", func(c *fiber.Ctx) error {
		return c.Download("./file/windwos_update_dev.bat")
	})

	app.Get("/update/ps", func(c *fiber.Ctx) error {
		return c.Download("./file/windwos_update_dev.ps1")
	})

	app.Get("/test2/:hostname/:winver", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"test":     "test",
			"ip":       c.IP(),
			"hostname": c.Params("hostname"),
			"winver":   c.Params("winver"),
		})
	})

	app.Get("/api/info_reg/:hostname/:winver", func(c *fiber.Ctx) error {
		// c.IP()
		// c.Params("hostname")
		// c.Params("winver")
		recivedata := c.JSON(fiber.Map{
			"ip":       c.IP(),
			"hostname": c.Params("hostname"),
			"winver":   c.Params("winver"),
		})

		return recivedata
	})

	app.Get("/api/db/insert", func(c *fiber.Ctx) error {

		return nil
	})

	app.Get("/", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	log.Fatal(app.Listen(":7979")) // http://localhost:7979/
}
