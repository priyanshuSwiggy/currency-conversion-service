package main

import (
	"currency-conversion-service/consumer"
	"currency-conversion-service/dao"
	"currency-conversion-service/server"
	"currency-conversion-service/util"
	"fmt"
	"log"
)

func main() {
	if err := util.LoadConfig("config.yaml"); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := dao.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("DynamoDB client initialized:", db)

	kafkaConsumer, err := consumer.NewKafkaConsumer()
	if err != nil {
		log.Fatal("Failed to create Kafka consumer:", err)
	}

	go server.StartGRPCServer()
	go consumer.ConsumeKafkaMessages(kafkaConsumer)
	server.StartHTTPServer()
}
