package middleware

import (
	"loan-app/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRateLimiterMiddleware(config *config.Config) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:               config.Security.RateLimit.MaxRequests,
		Expiration:        time.Duration(config.Security.RateLimit.Duration) * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"ok":    false,
				"error": "limiter/too-many-requests",
			})
		},
	})
}
