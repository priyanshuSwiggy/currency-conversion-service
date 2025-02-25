package main

import (
	"context"
	"currency-conversion-service/config"
	"currency-conversion-service/convert"
	"currency-conversion-service/money"
	pb "currency-conversion-service/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	from := money.Money{
		Currency: req.GetFrom().GetCurrency(),
		Amount:   req.GetFrom().GetAmount(),
	}
	toCurrency := req.ToCurrency

	converted, err := convert.Convert(from, toCurrency)
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

func main() {
	if err := config.LoadConversionRates("conversion_rates.json"); err != nil {
		log.Fatalf("Failed to load conversion rates: %v", err)
	}

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
