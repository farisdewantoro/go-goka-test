package controller

import (
	"errors"
	"fmt"
	"goka-test/internal/service/app"
	"goka-test/internal/service/dto"

	"github.com/kataras/iris/v12"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type TransactionController struct {
	app *app.Application
}

func NewTransactionController(app *app.Application) *TransactionController {
	return &TransactionController{
		app: app,
	}
}

func (w *TransactionController) DepositMoney(ctx iris.Context) {
	fmt.Println("executed 1")
	request := &dto.DepositMoneyRequest{}
	err := ctx.ReadJSON(request)
	if !errors.Is(err, nil) {
		ctx.JSON(Error{Code: 400, Message: "Bad Request"})
		return
	}
	fmt.Println("executed 2")
	w.app.DepositMoney.Handle(ctx, request)
	fmt.Println("executed 3")
}
