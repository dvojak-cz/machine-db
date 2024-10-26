package v1


import (
	"github.com/gofiber/fiber/v2"
)

type Api interface {
	app(*fiber.Ctx) error
	ping(*fiber.Ctx) error
}


func SetupRoutes(app *fiber.App, api Api) {
	app.Get("/ping", api.ping)
	app.Get("/app", api.app)
}
