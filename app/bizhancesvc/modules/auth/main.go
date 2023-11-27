package auth

import (
	"bizhancesvc/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthModule struct {
	dbPool       *pgxpool.Pool
	serverConfig configs.ServerConfig
}

var AuthModuleInstance AuthModule

func InitModule(dbPool *pgxpool.Pool, serverConfig configs.ServerConfig, router fiber.Router) {
	AuthModuleInstance = AuthModule{
		dbPool:       dbPool,
		serverConfig: serverConfig,
	}

	registerHandlers(router)
}
