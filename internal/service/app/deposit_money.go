package app

import (
	"context"
	"goka-test/internal/service/dto"
	kafkaMessages "goka-test/proto/kafka"
	"strconv"

	"github.com/lovoo/goka"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DepositMoneyHandler interface {
	Handle(ctx context.Context, request *dto.DepositMoneyRequest) error
}

type depositMoney struct {
	emitter *goka.Emitter
}

func NewDepositMoney(emitter *goka.Emitter) DepositMoneyHandler {
	return &depositMoney{
		emitter: emitter,
	}
}

func (d *depositMoney) Handle(ctx context.Context, request *dto.DepositMoneyRequest) error {
	m := kafkaMessages.Deposit{
		WalletId:  int64(request.WalletID),
		Amount:    int64(request.Amount),
		Timestamp: timestamppb.Now(),
	}

	return d.emitter.EmitSync(strconv.Itoa(request.WalletID), &m)
}
