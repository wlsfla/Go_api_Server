package api_router

import (
	"app/Lib/ConnManager"
	"app/Lib/api_router/Api_handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func init() {
	fmt.Println("[*] Apiv2Router init")
}

// API v2 Main Routing
func Apiv2Router() *fiber.App {
	app := fiber.New()

	app.Post("/insert/updatelog", Api_handler.Insertinfo)
	app.Get("/winver/:winver", Api_handler.GetBuildVer)
	app.Get("/file/:winver", ConnManager.Getupdatefile)

	app.Get("/connpool/info", ConnManager.GetConnInfo)
	app.Get("/connpool/max/:value", ConnManager.ChangeConnPool)

	// CurrConnPool info 가져오는지
	// MaxConnPool 변경
	// MaxConnPool 검사

	return app
}
