package v1


import (
	"github.com/gofiber/fiber/v2"
)

type Api interface {
	app(*fiber.Ctx) error
	ping(*fiber.Ctx) error
}

func SetupRoutesGroup(router fiber.Router, api Api) {
 	router.Get("/ping", api.ping)
 	router.Get("/app", api.app)
}
