package route

import "goka-test/internal/service/controller"

func (r RouteHandler) BalanceRouteHandler() {
	w := controller.NewWalletController(r.app)

	v1 := r.bootstrapper.Party("/details")
	{
		v1.Get("/{id:int}", w.GetBalanceDetail)
	}
}
