package route

func (a *Route) SetupUserRoute() {
	a.Log.Info("setting up user routes")

	a.App.Get("/v1/user/me", a.AuthMiddleware, a.UserHandler.FindSelf)
	a.Log.Info("mapped {/v1/user/me, GET} route")

	a.App.Get("/v1/user/limit", a.AuthMiddleware, a.UserHandler.FindLimits)
	a.Log.Info("mapped {/v1/user/limit, GET} route")
}
