package controller

import (
	"goka-test/internal/service/app"

	"github.com/kataras/iris/v12"
)

type WalletController struct {
	app *app.Application
}

func NewWalletController(app *app.Application) *WalletController {
	return &WalletController{
		app: app,
	}
}

func (w *WalletController) GetBalanceDetail(ctx iris.Context) {
	v, _ := w.app.GetBalanceWalletDetail.Handle(ctx)
	ctx.JSON(v)
}
