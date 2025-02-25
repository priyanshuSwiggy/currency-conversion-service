package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"currency-conversion-service/currency"
	pb "currency-conversion-service/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	fromCurrency, err := currency.GetCurrencyByCode(req.GetFromCurrency())
	if err != nil {
		return nil, err
	}
	toCurrency, err := currency.GetCurrencyByCode(req.GetToCurrency())
	if err != nil {
		return nil, err
	}
	amount := req.GetAmount()

	convertedAmount := currency.Convert(fromCurrency, toCurrency, amount)
	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &server{})

	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
