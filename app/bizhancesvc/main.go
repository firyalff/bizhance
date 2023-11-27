package main

import (
	"bizhancesvc/configs"
	"bizhancesvc/drivers"
	"context"
	"os"
	"strconv"

	"bizhancesvc/modules/auth"

	"github.com/mkideal/cli"
)

const AppVersion = "v0.0.1"

func main() {
	os.Exit(cli.Run(new(configs.ServerConfig), func(cliCtx *cli.Context) (err error) {
		ctx := context.Background()

		cfg := cliCtx.Argv().(*configs.ServerConfig)

		router := drivers.InitRouting()

		err = drivers.InitLogger(cfg.SentryDSN, router)
		if err != nil {
			panic(err)
		}

		DBPool, err := drivers.InitDBPool(ctx, cfg.DBURI)
		if err != nil {
			panic(err)
		}
		err = DBPool.Ping(ctx)
		if err != nil {
			panic(err)
		}

		defer DBPool.Close()

		baseRouter := router.Group("")

		registerHandlers(baseRouter)
		auth.InitModule(DBPool, *cfg, baseRouter)

		//routes below are protected by JWT

		return drivers.StartRouteServer(router, strconv.Itoa(cfg.ServerPort))
	}))
}
