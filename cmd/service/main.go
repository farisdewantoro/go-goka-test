package main

import (
	"goka-test/internal/service/config"
	"goka-test/internal/service/route"
	"goka-test/internal/service/service"
	"goka-test/pkg/http"
	"goka-test/pkg/kafka"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	kafka.LoadGlobalKafkaConfig(cfg.Kafka, []kafka.TopicConfig{
		{
			TopicName:  cfg.KafkaTopics.Deposit.TopicName,
			Partitions: cfg.KafkaTopics.Deposit.Partitions,
		},
	})

	app := newApp()
	app.Configure(Configure(cfg))
	app.Listen(":8100")
}

func newApp() *http.Bootstrapper {
	app := http.NewApplication()
	app.SetDefaultMiddleware().SetDefaultErrorHandlers()
	return app
}

func Configure(cfg *config.Config) http.Configurator {
	return func(b *http.Bootstrapper) {
		application := service.NewApplicationService(cfg, b)
		router := route.NewRouteHandler(b, application)
		router.RegisterRoutes()
	}

}
