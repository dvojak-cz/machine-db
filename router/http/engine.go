package httpEngine

import (
	v1Api "git.recombee.net/jtrojak/server-db/api/http/v1"
	"git.recombee.net/jtrojak/server-db/config"
	"git.recombee.net/jtrojak/server-db/constants"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func responce404(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
		"error": "Not found",
	})
}

func setCommonHeader(c *fiber.Ctx) error {
	c.Set("X-App-Name", constants.AppName)
	c.Set("X-App-Version", constants.Version)
	return c.Next()
}


func Run(cfg config.Http) {
	app := fiber.New(
		fiber.Config{
			AppName: constants.AppName,
			CaseSensitive: true,
			DisableDefaultContentType: false, //TODO
			DisableDefaultDate: false, //TODO
			DisableHeaderNormalizing: false, //TODO
			DisableKeepalive: false, //TODO
			DisablePreParseMultipartForm: false, //TODO
			DisableStartupMessage: true,
		})

	app.Use(setCommonHeader)
	app.Use(logger.New())
	app.Use(healthcheck.New())
	app.Use(requestid.New())

	api:= app.Group("/api")
	
	v1 := api.Group("/v1")
	v1Api.SetupRoutesGroup(v1, &v1Api.Base{})
	
	app.Use(responce404)
	epAddress := cfg.GetHttpAddress()
	log.Infof("Starting HTTP server on %s", epAddress)
	app.Listen(epAddress)
}
