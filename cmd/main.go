package main

import (
	"go-boilerplate/cmd/app"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	application, err := app.Wire()
	if err != nil {

		log.Fatalf("Failed auto injection to initialize application: %v", err)
	}

	application.Server.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "OK",
		})
	})

	if err := application.Server.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
