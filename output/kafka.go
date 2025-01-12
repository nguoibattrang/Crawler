package output

import (
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	return &KafkaProducer{
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic:   topic,
		}),
	}
}

func (kp *KafkaProducer) Produce(message []byte) error {
	return kp.Writer.WriteMessages(nil, kafka.Message{
		Value: message,
	})
}
