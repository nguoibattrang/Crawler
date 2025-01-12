package output

import (
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaProducer struct {
	Writer *kafka.Writer
	logger *zap.Logger
}

func NewKafkaProducer(brokers []string, topic string, logger *zap.Logger) *KafkaProducer {
	return &KafkaProducer{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		logger: logger,
	}
}

func (kp *KafkaProducer) Produce(mChan <-chan string) {
	for message := range mChan {
		err := kp.Writer.WriteMessages(nil, kafka.Message{
			Value: []byte(message),
		})
		if err != nil {
			kp.logger.Error("Failed to publish message", zap.Error(err))
		}
	}
}
