// kafkautil/consumer.go

package kafkautil

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(brokers []string, topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "test-group",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message from Kafka:", err)
			continue
		}

		log.Printf("Received message: Key=%s, Value=%s\n", string(msg.Key), string(msg.Value))
	}
}
