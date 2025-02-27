package main

import (
	"context"
	"currency-conversion-service/money"
	pb "currency-conversion-service/proto/moneyconverter"
	"currency-conversion-service/service"
	"currency-conversion-service/util"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMoneyConverterServer
	dsn string
}

type Config struct {
	DBConn       string
	KafkaBrokers string
	KafkaTopic   string
}

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

var config = Config{
	DBConn:       "host=localhost user=root password=root dbname=conversiondb port=5432 sslmode=disable",
	KafkaBrokers: "localhost:9092",
	KafkaTopic:   "currency_updates",
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	from := money.Money{
		Currency: req.GetFrom().GetCurrency(),
		Amount:   req.GetFrom().GetAmount(),
	}
	toCurrency := req.ToCurrency

	converted, err := service.ConvertMoney(s.dsn, from, toCurrency)
	if err != nil {
		return nil, err
	}
	return &pb.ConvertResponse{
		Converted: &pb.Money{
			Currency: converted.Currency,
			Amount:   converted.Amount,
		},
	}, nil
}

func startGRPCServer(dsn string) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyConverterServer(s, &server{dsn: dsn})
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startHTTPServer() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterMoneyConverterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	log.Println("HTTP server listening on :8085")
	http.ListenAndServe(":8085", mux)
}

func updateRateInDB(db *gorm.DB, rate ExchangeRate) error {
	return db.Table("conversion_rates").Where("currency = ?", rate.Currency).Updates(map[string]interface{}{"rate": rate.Rate}).Error
}

func consumeKafkaMessages(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
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
			var rate ExchangeRate
			if err := json.Unmarshal(msg.Value, &rate); err != nil {
				log.Println("Failed to parse Kafka message:", err)
				continue
			}
			if err := updateRateInDB(db, rate); err != nil {
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

func main() {
	if err := util.LoadConfig("config.yaml"); err != nil {
		log.Fatal("Failed to load config:", err)
	}
	dsn := "host=localhost user=root password=root dbname=conversiondb port=5432 sslmode=disable"

	go startGRPCServer(dsn)
	go consumeKafkaMessages(dsn)
	startHTTPServer()
}
