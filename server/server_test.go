package main

import (
	"context"
	"currency-conversion-service/money"
	pb "currency-conversion-service/pb"
	"currency-conversion-service/service"
	"currency-conversion-service/util"
	"testing"
)

type mockServer struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *mockServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	from := money.Money{
		Currency: req.GetFrom().GetCurrency(),
		Amount:   req.GetFrom().GetAmount(),
	}
	toCurrency := req.GetToCurrency()

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

func TestConvert(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"USD": 83.0,
		"EUR": 90.0,
	}

	server := &mockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "USD",
			Amount:   100.0,
		},
		ToCurrency: "INR",
	}

	res, err := server.Convert(context.Background(), req)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	expectedAmount := 8300.0
	if res.Converted.Amount != expectedAmount {
		t.Errorf("Expected %v, got %v", expectedAmount, res.Converted.Amount)
	}
}
