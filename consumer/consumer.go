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

type KafkaConsumer interface {
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
	Close() error
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error
}

type Database interface {
	UpdateRateInDB(rate dao.ExchangeRate) error
}

func ConsumeKafkaMessages(consumer KafkaConsumer) {
	db := &dao.DynamoDBClient{Client: dao.DynamoClient}
	ConsumeMessages(consumer, db)
	//for {
	//	msg, err := consumer.ReadMessage(-1)
	//	if err == nil {
	//		log.Printf("Received message: %s\n", string(msg.Value))
	//
	//		var rate dao.ExchangeRate
	//		if err := json.Unmarshal(msg.Value, &rate); err != nil {
	//			log.Println("Failed to parse Kafka message:", err)
	//			continue
	//		}
	//		if err := dao.UpdateRateInDB(rate); err != nil {
	//			log.Println("Failed to update database:", err)
	//		} else {
	//			log.Printf("Updated rate: %s = %f\n", rate.Currency, rate.Rate)
	//		}
	//	} else {
	//		log.Println("Error reading Kafka message:", err)
	//		time.Sleep(time.Second * 5)
	//	}
	//}
}

func ConsumeMessages(consumer KafkaConsumer, db Database) {
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil && msg != nil {
			log.Printf("Received message: %s\n", string(msg.Value))

			var rate dao.ExchangeRate
			if err := json.Unmarshal(msg.Value, &rate); err != nil {
				log.Println("Failed to parse Kafka message:", err)
				continue
			}
			if err := db.UpdateRateInDB(rate); err != nil {
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

func NewKafkaConsumer() (KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBrokers,
		"group.id":          "currency-conversion-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	consumer.SubscribeTopics([]string{config.KafkaTopic}, nil)
	return consumer, nil
}
