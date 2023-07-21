package main

import (
	"fmt"
	"profile/pkg/kafkautil"
	"profile/pkg/predicate"
	"sync"
)

const (
	brokerAddress = "localhost:9092"
	topic         = "test-topic"
	groupID       = "test-consumer-group"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Consumer goroutine
	go func() {
		defer wg.Done()
		kafkautil.ConsumeMessages([]string{brokerAddress}, topic)
	}()

	// Producer goroutine
	go func() {
		defer wg.Done()
		kafkautil.PublishMessage([]string{brokerAddress}, topic, "denemeler")
	}()

	wg.Wait()

	//log.Println("Starting..")
	//
	//// Kafka bağlantı parametreleri
	//brokers := []string{"localhost:9092"}
	//topic := "test-topic"
	//partition := 0
	//log.Println("DialLeader Start..")
	//
	//// Kafka admin istemcisi oluşturma
	//conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, partition)
	//if err != nil {
	//	log.Println("Error connecting to Kafka:", err)
	//	return
	//}
	//log.Println("DialLeader End..")
	//defer conn.Close()
	//
	//log.Println("kafka.CreateTopics..")
	//// Topic oluşturma
	//err = conn.CreateTopics(kafka.TopicConfig{
	//	Topic:             "test-topic2",
	//	NumPartitions:     1,
	//	ReplicationFactor: 1,
	//})
	//
	//if err != nil {
	//	log.Println("Error creating topic:", err)
	//	return
	//}
	//
	//log.Println("kafka.NewReader...")
	//// Kafka consumer ayarları
	//consumer := kafka.NewReader(kafka.ReaderConfig{
	//	Brokers:   brokers,
	//	Topic:     topic,
	//	GroupID:   "test-consumer-group",
	//	Partition: partition,
	//})
	//
	//log.Println("consumer.ReadMessage...")
	//// Kafka'dan mesajları tüketme
	//for {
	//	m, err := consumer.ReadMessage(context.Background())
	//	if err != nil {
	//		log.Println("Error reading message:", err)
	//		break
	//	}
	//
	//	log.Printf("Received message: key=%s, value=%s\n", string(m.Key), string(m.Value))
	//}
	//
	//log.Println("consumer.Close...")
	//// Consumer kapatma
	//err = consumer.Close()
	//if err != nil {
	//	log.Println("Error closing consumer:", err)
	//}

	expressionTree := &predicate.PredicateNode{
		Operator: "AND",
		Left: &predicate.PredicateNode{
			Operator: "OR",
			Left: &predicate.PredicateNode{
				Operator: ">",
				Left:     &predicate.ValueNode{Value: "age", ValueSpec: predicate.Column},
				Right:    &predicate.ValueNode{Value: 20},
			},
			Right: &predicate.PredicateNode{
				Operator: "AND",
				Left:     &predicate.ValueNode{Value: true},
				Right: &predicate.PredicateNode{
					Operator: "=",
					Left:     &predicate.ValueNode{Value: "name", ValueSpec: predicate.Column},
					Right:    &predicate.ValueNode{Value: "onur"},
				},
			},
		},
		Right: &predicate.PredicateNode{
			Operator: "OR",
			Left:     &predicate.ValueNode{Value: true},
			Right:    &predicate.ValueNode{Value: false},
		},
	}

	result := expressionTree.Evaluate()

	fmt.Println("Sonuç:", result)

	// e := echo.New()

	// e.POST("/profile/create", create)

	// e.POST("/users", func(c echo.Context) error {
	// 	user := new(User)
	// 	if err := c.Bind(user); err != nil {
	// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// 	}

	// 	// JSON objesini işleme
	// 	fmt.Println("Gelen Kullanıcı:", user)

	// 	return c.String(http.StatusOK, "Kullanıcı kaydedildi")
	// })

	// e.Logger.Fatal(e.Start(":1323"))

	/////////KAFKA
	//brokers := []string{"localhost:9092"}
	//topic := "example-topic"

	// Mesaj yayınlama
	//key := []byte("Key1")
	//value := []byte("Hello, Kafka!")
	//err := kafkautil.PublishMessage(brokers, topic, key, value)
	//if err != nil {
	// Hata durumunda işlemler
	//}

	// Mesajları tüketme
	//go kafkautil.ConsumeMessages(brokers, topic)

	// Main işlemleri devam eder...
}

// // e.POST("/save", save)
// func create(c echo.Context) error {
// 	// Get name and email
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	println("done")
// 	return c.String(http.StatusOK, "name:"+name+", email:"+email)
// }

//docker network create kafka-network
//
//docker run -d --name zookeeper --network kafka-network -e ZOOKEEPER_CLIENT_PORT=2181 -p 2181:2181 confluentinc/cp-zookeeper:6.2.0
//
//docker run -d --name kafka --network kafka-network -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092 -p 9092:9092 confluentinc/cp-kafka:6.2.0
