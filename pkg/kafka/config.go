package kafka

import (
	"log"

	"github.com/lovoo/goka"
)

// Config kafka config
type Config struct {
	Brokers           []string `mapstructure:"brokers"`
	GroupID           string   `mapstructure:"groupID"`
	InitialOffset     int64    `mapstructure:"initialOffset"`
	TableReplication  int      `mapstructure:"tableReplication"`
	StreamReplication int      `mapstructure:"streamReplication"`
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName  string `mapstructure:"topicName"`
	Partitions int    `mapstructure:"partitions"`
}

func LoadGlobalKafkaConfig(cfg *Config, topicCfg []TopicConfig) {
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = cfg.TableReplication
	tmc.Stream.Replication = cfg.StreamReplication

	config := goka.DefaultConfig()
	config.Consumer.Offsets.Initial = cfg.InitialOffset
	goka.ReplaceGlobalConfig(config)

	tm, err := goka.NewTopicManager(cfg.Brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}

	for _, v := range topicCfg {
		err = tm.EnsureStreamExists(string(v.TopicName), v.Partitions)
		if err != nil {
			log.Printf("Error creating kafka topic %s: %v", v.TopicName, err)
		}
	}
}
