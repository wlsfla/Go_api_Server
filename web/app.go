package main

import (
	db "app/DBConfig"
	handler "app/api_handler"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type hostinfo struct {
	host_ip   string
	host_name string
	winver    string
	build_ver string
	result    int
}

var server_ip string

func main() {
	defer db.Close()

	handler.Init()

	server_ip = "127.0.0.1"
	app := fiber.New()

	// accesslog write
	file, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// set log file
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time}\t|${ip}\t|${port}\t|${status}\t|${method}\t|${path}\n",
		TimeFormat: "2006-01-02|15:04:05",
		TimeZone:   "Asia/Seoul",
		Output:     io.MultiWriter(file, os.Stdout), // write file and stdout
	}))

	// ************************************************************************************
	// ************************************************************************************
	// refactoring

	app.Get("/cdn/chartjs", handler.Download_Chartjs) // http://localhost/cdn/chartjs
	app.Get("/", handler.Download_Downloader)

	api := app.Group("/api")
	v2 := api.Group("/v2")
	v2.Get("/buildver/?:winver", handler.GetBuildVer) //
	v2.Get("/file/:winver", handler.Download_updatefile)
	v2.Get("/updateinfo/:info", func(c *fiber.Ctx) error {
		// /updateinfo?
		// 		host_name=${host_name}&
		//		winver=${winver}&
		//		buildver=${buildver}&
		//		created_time=${created_time}&
		//		updated_time=${updated_time}&
		//		result=${result}

		hostip := c.IP()

		query := fmt.Sprintf("select host_ip from GoAPIService.update_info where host_ip = %s", hostip)

		info := hostinfo{
			hostip,
			c.Query("host_name"),
			c.Query("winver"),
			c.Query("buildver"),
			c.Query("result"),
		}

		return c.JSON(fiber.Map{
			"result":   1,
			"buildver": winverlist[c.Params("winver")],
		})
	})

	app.Get("/monitor", monitor.New(monitor.Config{Title: "Service Metrics Page", ChartJsURL: "http://" + server_ip + "/cdn/chartjs"}))
	log.Fatal(app.Listen(":9999")) // http://localhost:9999/
}
