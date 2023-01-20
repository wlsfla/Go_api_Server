package api_handler

import (
	db "app/DBConfig"
	model "app/model"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var Winverlist map[string]string

func Init() {
	// winverlist = getTarget_winver()
	init_winverlist()
}

func GetBuildVer(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"result":   1,
		"buildver": Winverlist[c.Params("winver")],
	})
}

func init_winverlist() {
	Winverlist = make(map[string]string)
	var winver, name string
	rows := db.Select("select * from GoAPIService.target_winver")

	for rows.Next() {
		err := rows.Scan(&winver, &name)
		if err != nil {
			log.Fatal(err)
		}

		Winverlist[winver] = name
	}
}

func InsertHostinfo(c *fiber.Ctx) error {
	info := model.New_Hostinfo(c)
	query := fmt.Sprintf("insert into GoAPIService.update_info values(%s, %s, %s, default, default, %s)", info.Host_ip, info.Host_name, info.Winver, info.Result)
	db.Update(query)

	return c.SendStatus(200)
}

func UpdateHostinfo(c *fiber.Ctx) error {
	// /insert/info/
	// 		buildver=${buildver}&
	// 		result=${result}

	info := model.New_Hostinfo(c)
	query := fmt.Sprintf("UPDATE GoAPIService.update_info SET result=%s  WHERE host_ip='%s'", info.Result, info.Host_ip)
	db.Update(query)

	return c.SendStatus(200)
}

func v2(api *fiber.Router) fiber.Router {
	// v2 := api.Group("/v2")

	return nil
}
