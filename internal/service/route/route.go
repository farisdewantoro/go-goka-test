package route

import (
	"goka-test/internal/service/app"
	"goka-test/pkg/http"
)

type RouteHandler struct {
	bootstrapper *http.Bootstrapper
	app          *app.Application
}

func NewRouteHandler(b *http.Bootstrapper, app *app.Application) RouteHandler {
	return RouteHandler{
		bootstrapper: b,
		app:          app,
	}
}

func (r *RouteHandler) RegisterRoutes() {
	r.BalanceRouteHandler()
	r.DepositRouteHandler()
}
