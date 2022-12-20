package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/mattn/go-sqlite3"
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

func dbinit() {
	db, err := sql.Open("sqlite3", "./log.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS 'update_log' ('host_ip' varchar(15) NOT NULL,'host_name' varchar(255) NOT NULL,'winver' char(5) NOT NULL,'result' int(5) NOT NULL,'created_time' datetime NOT NULL,'updated_time' datetime NOT NULL,'build_ver' varchar(255) NOT NULL,PRIMARY KEY ('host_ip'))")
	db.Exec("INSERT INTO 'update_log' VALUES ('1.1.1.1', 'host_name', '1901', 0, datetime('now', 'localtime'), datetime('now', 'localtime'), 1111.1111)") // dummy data

	rows, err := db.Query("SELECT * FROM update_log")
	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}

func insertData(info *hostinfo) {
	db, err := sql.Open("sqlite3", "./log.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	querystr := fmt.Sprintf("INSERT INTO 'update_log' VALUES ('%s', '%s', '%s', %d, datetime('now', 'localtime'), datetime('now', 'localtime'))", info.host_ip, info.host_name, info.winver, info.result)
	db.Exec(querystr)
}

func main() {
	dbinit()

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

	app.Get("/update", func(c *fiber.Ctx) error {
		return c.Download("./file/windwos_update_dev.bat")
	})

	app.Get("/update/ps", func(c *fiber.Ctx) error {
		return c.Download("./file/Scheduled_Register.ps1")
	})

	app.Get("/update/ps2", func(c *fiber.Ctx) error {
		return c.Download("./file/windwos_update_dev.ps1")
	})

	app.Get("/update/ps3", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/file/:winver", func(c *fiber.Ctx) error {
		return c.Download(fmt.Sprintf("./file/%s.msu", c.Params("winver")))
	})

	app.Get("/sysmon", func(c *fiber.Ctx) error {
		return c.Download("./file/sysmon.exe")
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
		db, err := sql.Open("sqlite3", "./log.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		query_str := fmt.Sprintf("update 'update_log' set result=%s, updated_time=datetime('now', 'localtime') where host_ip='%s'", c.Params("result"), c.IP())
		db.Exec(query_str)

		return c.JSON(fiber.Map{
			"ip":     c.IP(),
			"result": c.Params("result"),
			"query":  query_str,
		})
	})

	app.Get("/", monitor.New(monitor.Config{Title: "Service Metrics Page"}))

	log.Fatal(app.Listen(":9999")) // http://localhost:9999/
}
