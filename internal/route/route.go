package route

import (
	"loan-app/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Route struct {
	App            *fiber.App
	Log            *logrus.Logger
	AuthMiddleware fiber.Handler
	AuthHandler    *handler.AuthHandler
	UserHandler    *handler.UserHandler
}

func NewRoute(
	app *fiber.App,
	logger *logrus.Logger,
	authMiddleware fiber.Handler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
) *Route {
	return &Route{
		App:            app,
		Log:            logger,
		AuthMiddleware: authMiddleware,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
	}
}

func (a *Route) Setup() {
	a.Log.Info("setting up routes")

	a.SetupAuthRoute()
	a.SetupUserRoute()
	a.SetupSwaggerRoute()
}
