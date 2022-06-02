package service

import (
	"context"
	"fmt"
	"goka-test/internal/service/app"
	"goka-test/internal/service/config"
	"goka-test/pkg/dto"
	"goka-test/pkg/http"
	"log"

	"github.com/lovoo/goka"
)

func NewApplicationService(cfg *config.Config, b *http.Bootstrapper) *app.Application {
	emitter, err := goka.NewEmitter(cfg.Kafka.Brokers, goka.Stream(cfg.KafkaTopics.Deposit.TopicName), new(dto.DepositCodec))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	view, err := goka.NewView(cfg.Kafka.Brokers, goka.Table("aboveThreshold-table"), new(dto.AboveThresholdCodec))
	if err != nil {
		log.Fatalf("error creating view: %v", err)
	}
	go view.Run(context.Background())
	fmt.Println("vieww ", view)

	return &app.Application{
		DepositMoney:           app.NewDepositMoney(emitter),
		GetBalanceWalletDetail: app.NewGetBalanceWalletDetail(view),
	}
}
