package route

import "goka-test/internal/service/controller"

func (r RouteHandler) DepositRouteHandler() {
	c := controller.NewTransactionController(r.app)
	v1 := r.bootstrapper.Party("/deposits")
	{
		v1.Post("/", c.DepositMoney)
	}
}
