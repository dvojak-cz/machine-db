package httpEngine

import (
	"git.recombee.net/jtrojak/server-db/config"
	v1Api "git.recombee.net/jtrojak/server-db/api/http/v1"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Run(cfg config.Http) {
	app := fiber.New()
	epAddress := cfg.GetHttpAddress()
	v1Api.SetupRoutes(app, &v1Api.Base{})
	log.Infof("Starting HTTP server on %s", epAddress)
	app.Listen(epAddress)
}
