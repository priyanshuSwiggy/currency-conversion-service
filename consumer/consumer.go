package consumer

import (
	"currency-conversion-service/dao"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

type Config struct {
	KafkaBrokers string
	KafkaTopic   string
}

var config = Config{
	KafkaBrokers: "localhost:9092",
	KafkaTopic:   "currency_updates",
}

func ConsumeKafkaMessages() {
	_, err := dao.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBrokers,
		"group.id":          "currency-conversion-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	defer consumer.Close()
	consumer.SubscribeTopics([]string{config.KafkaTopic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))

			var rate dao.ExchangeRate
			if err := json.Unmarshal(msg.Value, &rate); err != nil {
				log.Println("Failed to parse Kafka message:", err)
				continue
			}
			if err := dao.UpdateRateInDB(rate); err != nil {
				log.Println("Failed to update database:", err)
			} else {
				log.Printf("Updated rate: %s = %f\n", rate.Currency, rate.Rate)
			}
		} else {
			log.Println("Error reading Kafka message:", err)
			time.Sleep(time.Second * 5)
		}
	}
}
