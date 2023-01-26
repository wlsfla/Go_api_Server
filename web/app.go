package main

import (
	db "app/DBConfig"
	api_handler "app/api_handler"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var server_ip string

func main() {
	db.Init()
	api_handler.Init()

	server_ip = "127.0.0.1"

	// accesslog write
	file, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// ************************************************************************************
	// refactoring

	// app.Get("/", api_handler.Download_UpdateDownloader)
	// app.Get("/file/:winver", api_handler.Download_updatefile)

	// api := app.Group("/api")
	// v2 := api.Group("/v2")
	// SetRoutes(&v2)

	app := fiber.New()
	app.Mount("/api/v2", api_handler.Apiv2())

	SetStaticAsset(app)

	app.Get("/", monitor.New(monitor.Config{Title: "Service Metrics Page", ChartJsURL: "http://" + server_ip + "/static/js/Chart.bundle.min.js"}))

	// set log file
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time}\t${ip}\t${status}\t${method}\t${path}\n",
		TimeFormat: "2006-01-02|15:04:05",
		TimeZone:   "Asia/Seoul",
		Output:     io.MultiWriter(file, os.Stdout), // write file and stdout
	}))

	log.Fatal(app.Listen(":9999")) // http://localhost:9999/
}

func SetStaticAsset(app *fiber.App) {
	app.Static("/static", "./static")
	/*
		/static/files/test.txt
		/static/js/Chart.bundle.min.js

		/static/files/win_update.zip
	*/
}
