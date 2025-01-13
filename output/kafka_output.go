package output

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nguoibattrang/crawler/crawl"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaProducer struct {
	writer *kafka.Writer
	log    *zap.Logger
}

func NewKafkaProducer(brokers []string, topic string, log *zap.Logger) (*KafkaProducer, error) {
	err := ensureTopicExists(brokers, topic, log)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		log: log,
	}, nil
}

func (inst *KafkaProducer) Produce(mChan <-chan crawl.Data) {
	for message := range mChan {
		m, err := json.Marshal(message)
		if m != nil {
			inst.log.Error("Convert data to JSON fail", zap.Error(err))
			continue
		}
		err = inst.writer.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(m),
		})
		if err != nil {
			inst.log.Panic("Failed to publish message", zap.Error(err))
		}
	}
}

func ensureTopicExists(brokers []string, topic string, log *zap.Logger) error {
	// Create a Kafka connection to the broker
	conn, err := kafka.Dial("tcp", kafka.TCP(brokers...).String())
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka broker: %w", err)
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     -1,
			ReplicationFactor: -1,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Error("failed to create topic", zap.Error(err))
		return err
	}

	log.Info("Topic created", zap.String("topic", topic))
	return nil
}
