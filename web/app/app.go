package main

import (
	db "app/app/DBConfig"
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
var winverlist map[string]string

func main() {
	defer db.Close()

	winverlist = getTarget_winver()
	server_ip = "127.0.0.1"
	app := fiber.New()

	// accesslog write
	file, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Download("./file/win_update.zip")
	})

	app.Get("/update/ps3", func(c *fiber.Ctx) error {
		return c.SendString("OK") // 스케줄러에서 요청받는 부분
	})

	app.Get("/file/:winver", func(c *fiber.Ctx) error {
		return c.Download(fmt.Sprintf("./file/updatefile/%s.msu", c.Params("winver")))
	})

	app.Get("/file/:winver", func(c *fiber.Ctx) error {
		return c.Download(fmt.Sprintf("./file/%s.msu", c.Params("winver")))
	})

	app.Get("/sysmon", func(c *fiber.Ctx) error {
		return c.Download("./file/sysmon.exe")
	})

	app.Get("/cdn/chartjs", func(c *fiber.Ctx) error { // http://localhost/cdn/chartjs
		return c.Download("./js/Chart.bundle.min.js")
	})

	app.Get("/api/info_reg/:hostname/:winver/:build", func(c *fiber.Ctx) error {

		info := hostinfo{
			c.IP(),
			c.Params("hostname"),
			c.Params("winver"),
			c.Params("build"),
			0,
		}

		query := fmt.Sprintf("insert into GoAPIService.update_info values ('%s', '%s', '%s', '%s', default, default, %d)", info.host_ip, info.host_name, info.winver, info.build_ver, info.result)
		db.Insert(query)

		recivedata := c.JSON(fiber.Map{
			"ip":       c.IP(),
			"hostname": c.Params("hostname"),
			"winver":   c.Params("winver"),
			"build":    c.Params("build"),
		})

		return recivedata
	})

	app.Get("/api/result/:result", func(c *fiber.Ctx) error {
		query := fmt.Sprintf("UPDATE GoAPIService.update_info SET result=%s  WHERE host_ip='%s'", c.Params("result"), c.IP())
		db.Update(query)

		return c.JSON(fiber.Map{
			"ip":     c.IP(),
			"result": c.Params("result"),
			"query":  query,
		})
	})

	app.Get("/winver/:winver", func(c *fiber.Ctx) error {

		return c.SendString(
			winverlist[c.Params("winver")],
		)
	})

	// ************************************************************************************
	// refactoring
	api := app.Group("/api")
	v2 := api.Group("/v2")
	v2.Get("/winver/:winver", getwinver)
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
