package route

import (
	"loan-app/internal/middleware"
	"loan-app/internal/model"
)

func (a *Route) SetupAssetRoute() {
	a.Log.Info("setting up asset routes")

	a.App.Post("/v1/asset", a.AuthMiddleware, middleware.RoleMiddleware(model.RoleAdmin), a.AssetHandler.Create)
	a.Log.Info("mapped {/v1/asset, POST} route")
	a.App.Get("/v1/asset", a.AuthMiddleware, a.AssetHandler.Lists)
	a.Log.Info("mapped {/v1/asset, GET} route")
}
