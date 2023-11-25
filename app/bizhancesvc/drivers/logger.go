package drivers

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func InitLogger(sentryDSN string, serverRouter *fiber.App) (err error) {
	err = sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if c, ok := hint.Context.Value(sentry.RequestContextKey).(*fiber.Ctx); ok {
					// You have access to the original Context if it panicked
					fmt.Println(utils.ImmutableString(c.Hostname()))
				}
			}
			fmt.Println(event)
			return event
		},
		EnableTracing: false,
	})
	if err != nil {
		return err
	}

	serverRouter.Use(fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: true,
	}))

	return
}
