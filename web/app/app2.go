package main

import (
	db "app/app/DBConfig"
	"log"

	"github.com/gofiber/fiber/v2"
)

func getwinver(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"result":   1,
		"buildver": winverlist[c.Params("winver")],
	})
}

func getTarget_winver() map[string]string {
	winverlist := make(map[string]string)
	var winver, name string
	rows := db.Select("select * from GoAPIService.target_winver")

	for rows.Next() {
		err := rows.Scan(&winver, &name)
		if err != nil {
			log.Fatal(err)
		}

		winverlist[winver] = name
	}
	return winverlist
}

func insertUpdateinfo() {

}
