package route

import (
	"loan-app/internal/middleware"
	"loan-app/internal/model"
)

func (a *Route) SetupTransactionRoute() {
	a.Log.Info("setting up transaction routes")

	a.App.Post("/v1/transaction", a.AuthMiddleware, middleware.RoleMiddleware(model.RoleCustomer), a.TransactionHandler.Create)
	a.Log.Info("mapped {/v1/transaction, POST} route")
	a.App.Get("/v1/transaction", a.AuthMiddleware, middleware.RoleMiddleware(model.RoleCustomer), a.TransactionHandler.Lists)
	a.Log.Info("mapped {/v1/transaction, GET} route")
}
