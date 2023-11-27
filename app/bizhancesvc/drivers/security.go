package drivers

import (
	jwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func InitJWT(jwtSecret string, serverRouter *fiber.App) {
	serverRouter.Use(jwt.New(jwt.Config{
		SigningKey: jwt.SigningKey{Key: []byte(jwtSecret)},
	}))
}
