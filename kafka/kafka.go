package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/massigerardi/trader-slack-notifier-go/model"
	"github.com/massigerardi/trader-slack-notifier-go/services"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	kafkaBrokerUrl     string
	kafkaVerbose       bool
	kafkaTopic         string
	kafkaConsumerGroup string
	kafkaClientId      string
	notifier           services.Notifier
}

func NewConsumer(kafkaBrokerUrl, kafkaTopic, kafkaConsumerGroup, kafkaClientId string, kafkaVerbose bool, notifier services.Notifier) Consumer {
	return Consumer{
		kafkaBrokerUrl:     kafkaBrokerUrl,
		kafkaVerbose:       kafkaVerbose,
		kafkaTopic:         kafkaTopic,
		kafkaConsumerGroup: kafkaConsumerGroup,
		kafkaClientId:      kafkaClientId,
		notifier:           notifier,
	}
}

func (consumer *Consumer) Run() {

	brokers := strings.Split(consumer.kafkaBrokerUrl, ",")

	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         consumer.kafkaClientId,
		Topic:           consumer.kafkaTopic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()
	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Msgf("error while receiving message: %s", err.Error())
			continue
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		message := &model.MessageRequest{}
		if err := json.Unmarshal(m.Value, message); err != nil {
			log.Error().Msgf("error for %s: %s", string(m.Value), err.Error())
			continue
		}
		consumer.notifier.SendMessage(nil, *message)
	}
}
