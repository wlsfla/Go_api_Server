package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type hostinfo struct {
	host_ip      string
	host_name    string
	winver       string
	result       int
	created_time string
	updated_time string
	build_ver    string
}

var server_ip string

func insertData(info *hostinfo) {
	// insert into GoAPIService.pc_info values ('1.1.1.1', 'dummy_host', '1900', '1111.1111', default, default, 0);
	db, _ := sql.Open("mysql", "root:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService")
	defer db.Close()

	querystr := fmt.Sprintf("insert into GoAPIService.pc_info values ('%s', '%s', '%s', '%s', default, default, 0)", info.host_ip, info.host_name, info.winver, info.build_ver)
	db.Exec(querystr)
}

func main() {
	server_ip = "127.0.0.1"
	app := fiber.New()

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
			0,
			"",
			"",
			c.Params("build"),
		}

		insertData(&info)

		recivedata := c.JSON(fiber.Map{
			"ip":       c.IP(),
			"hostname": c.Params("hostname"),
			"winver":   c.Params("winver"),
			"build":    c.Params("build"),
		})

		return recivedata
	})

	app.Get("/api/result/:result", func(c *fiber.Ctx) error {
		// winupdate 결과 table update
		db, _ := sql.Open("mysql", "root:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService")
		defer db.Close()

		// UPDATE GoAPIService.pc_info SET updated_time = NOW() WHERE host_ip = '1.1.1.1';
		// query_str := fmt.Sprintf("update 'update_log' set result=%s, updated_time=datetime('now', 'localtime') where host_ip='%s'", c.Params("result"), c.IP())
		query_str := fmt.Sprintf("UPDATE GoAPIService.pc_info SET result=%s, updated_time=NOW() WHERE host_ip='%s'", c.Params("result"), c.IP())
		db.Exec(query_str)

		return c.JSON(fiber.Map{
			"ip":     c.IP(),
			"result": c.Params("result"),
			"query":  query_str,
		})
	})

	dbconnTest()
	app.Get("/api/monitor", monitor.New(monitor.Config{Title: "Service Metrics Page", ChartJsURL: "http://" + server_ip + "/cdn/chartjs"}))
	log.Fatal(app.Listen(":9999")) // http://localhost:9999/
}

func dbconnTest() {
	// if using mariaDB on docker-container, You have to change ip address "host.docker.internal"
	// issue https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host
	db, err := sql.Open("mysql", "root:1q2w3e4r!@tcp(host.docker.internal:3306)/GoAPIService")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
