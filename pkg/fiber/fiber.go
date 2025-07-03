package fiber

import (
	"errors"
	"fmt"
	"loan-app/config"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewFiber(config *config.Config, log *logrus.Logger) *fiber.App {
	// Create context logger for error handler

	var app = fiber.New(fiber.Config{
		AppName:               fmt.Sprintf("%s - %s", config.App.Name, config.App.Version),
		ServerHeader:          "loan-app",
		ReadTimeout:           time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout:          time.Duration(config.App.WriteTimeout) * time.Second,
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler(log),
		Prefork:               config.App.Prefork,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	})

	return app
}

func errorHandler(log *logrus.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		log.Errorf("request error occurred: %s", err)

		return ctx.Status(code).JSON(fiber.Map{
			"ok":     false,
			"errors": "internal/server-error",
		})
	}
}
