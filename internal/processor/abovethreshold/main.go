package abovethreshold

import (
	"context"
	"fmt"
	"goka-test/internal/processor/abovethreshold/config"
	"goka-test/pkg/dto"
	kafkaMessages "goka-test/proto/kafka"
	"log"

	"github.com/lovoo/goka"
)

type AboveThresholdProcessor interface {
	Run(ctx context.Context) error
}

type aboveThresholdProcessor struct {
	cnf *config.Config
}

func NewAboveThresholdProcessor(cnf *config.Config) AboveThresholdProcessor {
	return &aboveThresholdProcessor{
		cnf: cnf,
	}
}

func (a aboveThresholdProcessor) rollingWindowThreshold(ctx goka.Context, msg interface{}) {
	var aboveThreshold *kafkaMessages.AboveThreshold
	if v := ctx.Value(); v != nil {
		fmt.Println("hereeeeee ", v)
		aboveThreshold = v.(*kafkaMessages.AboveThreshold)
	}
	log.Println("aboveThreshold ", aboveThreshold)

	m := msg.(*kafkaMessages.Deposit)

	log.Println(m)

	if aboveThreshold == nil {
		aboveThreshold = &kafkaMessages.AboveThreshold{
			Deposits:         []*kafkaMessages.Deposit{m},
			IsAboveThreshold: false,
		}
	} else {
		fmt.Println("here")
		aboveThreshold.Deposits = append(aboveThreshold.Deposits, m)
	}

	// // if len(aboveThreshold.Deposits) > maxMessages {
	// // 	ml = ml[len(ml)-maxMessages:]
	// // }
	ctx.SetValue(aboveThreshold)
}

func (a aboveThresholdProcessor) Run(ctx context.Context) error {
	g := goka.DefineGroup(goka.Group(a.cnf.Kafka.GroupID),
		goka.Input(goka.Stream(a.cnf.KafkaTopics.Deposit.TopicName), new(dto.DepositCodec), a.rollingWindowThreshold),
		goka.Persist(new(dto.AboveThresholdCodec)),
	)
	p, err := goka.NewProcessor(a.cnf.Kafka.Brokers, g)
	if err != nil {
		return err
	}
	return p.Run(ctx)
}
