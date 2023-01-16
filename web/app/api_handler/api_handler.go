package api_handler

import (
	db "app/app/DBConfig"
	"log"

	"github.com/gofiber/fiber/v2"
)

var winverlist map[string]string

func Init() {
	// winverlist = getTarget_winver()
	init_winverlist()
}

func Getwinver(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"result":   1,
		"buildver": winverlist[c.Params("winver")],
	})
}

func init_winverlist() {
	winverlist = make(map[string]string)
	var winver, name string
	rows := db.Select("select * from GoAPIService.target_winver")

	for rows.Next() {
		err := rows.Scan(&winver, &name)
		if err != nil {
			log.Fatal(err)
		}

		winverlist[winver] = name
	}
}

func insertUpdateinfo() {

}
