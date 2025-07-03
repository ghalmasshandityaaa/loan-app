package middleware

import (
	"loan-app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCorsMiddleware(config *config.Config) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     config.Security.Cors.AllowedOrigins,
		AllowMethods:     config.Security.Cors.AllowedMethods,
		AllowCredentials: config.Security.Cors.AllowCredentials,
	})
}
