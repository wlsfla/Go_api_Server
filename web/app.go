package main

import (
	// _ "app/LibDBConfig"
	"app/Lib/Logging"
	"app/Lib/api_router"
	"app/Lib/common"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {

	// ************************************************************************************
	// refactoring

	// app.Get("/", api_handler.Download_UpdateDownloader)
	// app.Get("/file/:winver", api_handler.Download_updatefile)

	// api := app.Group("/api")
	// v2 := api.Group("/v2")
	// SetRoutes(&v2)

	app := fiber.New()
	SetStaticAsset(app)
	app.Mount("/api/v2", api_router.Apiv2Router())
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Service Metrics Page", ChartJsURL: "http://" + common.Server_ip + "/static/js/Chart.bundle.min.js"}))

	// set log file
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time}\t${ip}\t${status}\t${method}\t${path}\n",
		TimeFormat: "2006-01-02|15:04:05",
		TimeZone:   "Asia/Seoul",
		Output:     io.MultiWriter(Logging.Fileptr, os.Stdout), // write file and stdout
	}))
	defer Logging.Close()

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
