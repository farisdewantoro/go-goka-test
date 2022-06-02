package config

import (
	"flag"
	"fmt"
	kafkaClient "goka-test/pkg/kafka"
	"os"

	"goka-test/pkg/constant"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "config path")
}

type Config struct {
	Kafka       *kafkaClient.Config `mapstructure:"kafka"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
}

type KafkaTopics struct {
	Deposit kafkaClient.TopicConfig `mapstructure:"deposit"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constant.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/internal/service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constant.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	kafkaBrokers := os.Getenv(constant.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
