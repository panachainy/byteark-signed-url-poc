package app

import (
	"go-boilerplate/cmd/config"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	app     = &Application{}
	appOnce sync.Once
)

type Application struct {
	Server *fiber.App
	Config *config.Configuration
	Log    *logrus.Logger
}

func Provide(
	config *config.Configuration,
	log *logrus.Logger,
) *Application {
	appOnce.Do(func() {
		app = &Application{
			Config: config,
			Server: fiber.New(
				fiber.Config{
					// Global exception handler.
					ErrorHandler: func(ctx *fiber.Ctx, err error) error {
						code := fiber.StatusInternalServerError

						if e, ok := err.(*fiber.Error); ok {
							code = e.Code
						}

						return ctx.
							Status(code).
							JSON(fiber.Map{
								"message": err.Error(),
								"ok":      false,
							})
					},
				},
			),
			Log: log,
		}
	})

	return app
}

// Only for test
func Reset() {
	appOnce = sync.Once{}
}
