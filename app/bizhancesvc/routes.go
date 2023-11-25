package main

import (
	"bizhancesvc/configs"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(router fiber.Router) {
	router.Get("/version", getVersionHandler)
}

func getVersionHandler(ctx *fiber.Ctx) error {
	type appVer struct {
		Version string `json:"version"`
	}
	return ctx.JSON(appVer{
		Version: configs.AppVersion,
	})
}
