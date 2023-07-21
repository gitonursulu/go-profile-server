// kafkautil/publisher.go

package kafkautil

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func PublishMessage(brokers []string, topic string, data string) error {
	//writer := &kafka.Writer{
	//	Addr:     kafka.TCP(brokers...),
	//	Topic:    topic,
	//	Balancer: &kafka.LeastBytes{},
	//}

	// Kafka bağlantı parametreleri
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		log.Fatalf("Error connecting to Kafka: %v", err)
	}
	defer conn.Close()

	// Mesaj gönderme döngüsü
	for i := 1; i <= 5; i++ {
		message := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", data+"_"+strconv.Itoa(i))),
			Value: []byte(fmt.Sprintf("Message-%d", data+"_"+strconv.Itoa(i))),
		}

		_, err := conn.WriteMessages(message)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}

		log.Printf("Producer: Sent message: key=%s, value=%s\n", string(message.Key), string(message.Value))

		time.Sleep(1 * time.Second)
	}

	//err := writer.WriteMessages(context.Background(), message)
	//if err != nil {
	//	log.Println("Error publishing message to Kafka:", err)
	//}
	//
	//err = writer.Close()
	//if err != nil {
	//	log.Println("Error closing Kafka writer:", err)
	//}

	return err
}
