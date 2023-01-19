package api_handler

import (
	db "app/DBConfig"
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

func insertUpdateinfo() {

}
