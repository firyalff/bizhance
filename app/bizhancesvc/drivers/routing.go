package drivers

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPMethod int32

const (
	HTTPGET HTTPMethod = iota
	HTTPPOST
	HTTPPUT
	HTTPDELETE
)

type AppRoute struct {
	Path    string
	Method  HTTPMethod
	Handler func(ctx *fiber.Ctx) error
}

func InitRouting() (router *fiber.App) {
	router = fiber.New()

	return router
}

func StartRouteServer(router *fiber.App, port string) (err error) {
	return router.Listen(":" + port)
}

func RegisterRoutes(router *fiber.App, appRoutes []AppRoute) {
	for _, appRoute := range appRoutes {
		switch appRoute.Method {
		case HTTPGET:
			router.Get(appRoute.Path, appRoute.Handler)
		}
	}
}
