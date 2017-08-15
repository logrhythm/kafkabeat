package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher/bc/publisher"

	"github.com/Shopify/sarama"
	"github.com/justsocialapps/kafkabeat/config"
	cluster "gopkg.in/bsm/sarama-cluster.v2"
)

type Kafkabeat struct {
	done     chan struct{}
	config   config.Config
	client   publisher.Client
	consumer *cluster.Consumer
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	kafkaConfig := cluster.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Group.Return.Notifications = true
	kafkaConfig.Config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := cluster.NewConsumer(config.Brokers, config.Group, config.Topics, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("Error starting consumer on %s: %s", "tracking", err)
	}

	bt := &Kafkabeat{
		done:     make(chan struct{}),
		config:   config,
		consumer: consumer,
	}
	return bt, nil
}

func (bt *Kafkabeat) Run(b *beat.Beat) error {
	logp.Info("kafkabeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()

	for {
		select {
		case <-bt.done:
			bt.consumer.Close()
			return nil
		case ev := <-bt.consumer.Messages():
			event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       b.Info.Name,
				"message":    string(ev.Value),
			}
			if bt.client.PublishEvent(event, publisher.Guaranteed, publisher.Sync) {
				bt.consumer.MarkOffset(ev, "")
			}
		case notification := <-bt.consumer.Notifications():
			logp.Info("Rebalanced: %+v", notification)
		case err := <-bt.consumer.Errors():
			logp.Err("Error in Kafka consumer: %s", err.Error())
		}
	}
}

func (bt *Kafkabeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
