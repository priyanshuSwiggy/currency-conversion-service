package main

import (
	"context"
	"currency-conversion-service/money"
	pb "currency-conversion-service/proto/moneyconverter"
	"currency-conversion-service/service"
	"currency-conversion-service/util"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMoneyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	from := money.Money{
		Currency: req.GetFrom().GetCurrency(),
		Amount:   req.GetFrom().GetAmount(),
	}
	toCurrency := req.ToCurrency

	converted, err := service.ConvertMoney(from, toCurrency)
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

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyConverterServer(s, &server{})
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
	if err := util.LoadConversionRates("conversion_rates.json"); err != nil {
		log.Fatalf("Failed to load conversion rates: %v", err)
	}

	go startGRPCServer()
	startHTTPServer()
}
