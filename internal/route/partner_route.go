package route

import (
	"loan-app/internal/middleware"
	"loan-app/internal/model"
)

func (a *Route) SetupPartnerRoute() {
	a.Log.Info("setting up partner routes")

	a.App.Post("/v1/partner", a.AuthMiddleware, middleware.RoleMiddleware(model.RoleAdmin), a.PartnerHandler.Create)
	a.Log.Info("mapped {/v1/partner, POST} route")
	a.App.Get("/v1/partner", a.AuthMiddleware, a.PartnerHandler.Lists)
	a.Log.Info("mapped {/v1/partner, GET} route")
}
