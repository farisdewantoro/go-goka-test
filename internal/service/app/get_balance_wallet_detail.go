package app

import (
	"context"
	kafkaMessages "goka-test/proto/kafka"

	"github.com/lovoo/goka"
)

type GetBalanceWalletDetailHandler interface {
	Handle(ctx context.Context) (*kafkaMessages.AboveThreshold, error)
}

type getBalanceWalletDetail struct {
	view *goka.View
}

func NewGetBalanceWalletDetail(view *goka.View) GetBalanceWalletDetailHandler {
	return &getBalanceWalletDetail{
		view: view,
	}
}

func (g *getBalanceWalletDetail) Handle(ctx context.Context) (*kafkaMessages.AboveThreshold, error) {
	v, err := g.view.Get("1000")

	if err != nil {
		return nil, err
	}

	result, _ := v.(*kafkaMessages.AboveThreshold)

	return result, nil
}
