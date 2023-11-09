package main

import (
	"fmt"
	"go-boilerplate/cmd/app"
	"go-boilerplate/cmd/config"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	bytearkSignedURL "github.com/byteark/byteark-sdk-go"
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

	application.Server.Get("/t", func(ctx *fiber.Ctx) error {

		var signedUrl = byteark(application.Config)
		return ctx.JSON(fiber.Map{
			"status":    "OK",
			"signedUrl": signedUrl,
		})
	})

	if err := application.Server.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

// POC signed url
func byteark(config *config.Configuration) string {
	signerOptions := bytearkSignedURL.SignerOptions{
		AccessID:     config.BytearkAccessId,
		AccessSecret: config.BytearkAccessSecret,
	}

	createSignerErr := bytearkSignedURL.CreateSigner(signerOptions)

	if createSignerErr != nil {
		panic(createSignerErr)
	}

	signOptions := bytearkSignedURL.SignOptions{
		// "path_prefix": "/live/",
	}

	signedURL, signErr := bytearkSignedURL.Sign(
		// FIXME: get hls url from db.
		config.BytearkHlsUrl,
		// Note: (unix second) this is timestamp, u can ref from https://www.epochconverter.com
		1702115115,
		signOptions,
	)

	if signErr != nil {
		panic(signErr)
	}

	// note: 0 is trigger to use default in library `int(time.Now().Unix())`
	isPassed, verifyErr := bytearkSignedURL.Verify(signedURL, 0)
	if verifyErr != nil {
		panic(verifyErr)
	}

	if !isPassed {
		msg := fmt.Sprintf("\n verify not pass  %s", strconv.FormatBool(isPassed))
		panic(msg)
	}

	return signedURL
}
