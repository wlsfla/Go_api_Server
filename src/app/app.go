package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	file, err := os.OpenFile("result.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time}\t${pid}\t${locals:requestid}\t${status}\t${method} ${path}​\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Seoul",
		Output:     file,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	log.Fatal(app.Listen(":7979"))
}
