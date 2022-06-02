package main

import (
	"context"
	"fmt"
	"goka-test/internal/processor/abovethreshold"
	"goka-test/internal/processor/abovethreshold/config"
	"goka-test/pkg/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.Kafka)
	kafka.LoadGlobalKafkaConfig(cfg.Kafka, []kafka.TopicConfig{
		{
			TopicName:  cfg.KafkaTopics.Deposit.TopicName,
			Partitions: cfg.KafkaTopics.Deposit.Partitions,
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	processor := abovethreshold.NewAboveThresholdProcessor(cfg)
	log.Println("starting to run above-processor")
	go func() {
		defer close(done)
		if err = processor.Run(ctx); err != nil {
			log.Printf("error running processor: %v", err)
		}
	}()
	// Wait for SIGINT/SIGTERM
	sigs := make(chan os.Signal)
	go func() {
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	}()

	select {
	case <-sigs:
	case <-done:
	}
	cancel()
	<-done
	log.Println("done")
}
