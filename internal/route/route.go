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
	PartnerHandler *handler.PartnerHandler
}

func NewRoute(
	app *fiber.App,
	logger *logrus.Logger,
	authMiddleware fiber.Handler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	partnerHandler *handler.PartnerHandler,
) *Route {
	return &Route{
		App:            app,
		Log:            logger,
		AuthMiddleware: authMiddleware,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		PartnerHandler: partnerHandler,
	}
}

func (a *Route) Setup() {
	a.Log.Info("setting up routes")

	a.SetupAuthRoute()
	a.SetupUserRoute()
	a.SetupPartnerRoute()
	a.SetupSwaggerRoute()
}
