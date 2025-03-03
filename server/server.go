package main

import (
	"context"
	"currency-conversion-service/consumer"
	"currency-conversion-service/money"
	pb "currency-conversion-service/proto/moneyconverter"
	"currency-conversion-service/service"
	"currency-conversion-service/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type server struct {
	pb.UnimplementedMoneyConverterServer
	dsn string
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

func main() {
	if err := util.LoadConfig("config.yaml"); err != nil {
		log.Fatal("Failed to load config:", err)
	}
	dsn := "host=localhost user=root password=root dbname=conversiondb port=5432 sslmode=disable"

	go startGRPCServer(dsn)
	go consumer.ConsumeKafkaMessages(dsn)
	startHTTPServer()
}
