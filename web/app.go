package main

import (
	"app/DBConfig"
	"app/api_handler"
	"app/common"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var server_ip string

func main() {
	DBConfig.Init()
	api_handler.Init()

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
	app.Mount("/api/v2", api_handler.Apiv2Router())

	SetStaticAsset(app)

	app.Get("/monitor", monitor.New(monitor.Config{Title: "Service Metrics Page", ChartJsURL: "http://" + common.Server_ip + "/static/js/Chart.bundle.min.js"}))

	// set log file
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time}\t${ip}\t${status}\t${method}\t${path}\n",
		TimeFormat: "2006-01-02|15:04:05",
		TimeZone:   "Asia/Seoul",
		Output:     io.MultiWriter(file, os.Stdout), // write file and stdout
	}))

	log.Fatal(app.Listen(":9999"))
}

func SetStaticAsset(app *fiber.App) {
	app.Static("/static", "./static")
	/*
		/static/js/Chart.bundle.min.js
		/static/files/21H2.msu
		/static/files/win_update.zip

	*/
}
