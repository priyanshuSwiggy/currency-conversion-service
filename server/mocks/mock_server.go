package mocks

import (
	"context"
	"currency-conversion-service/money"
	pb "currency-conversion-service/pb"
	"currency-conversion-service/service"
)

type MockServer struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *MockServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
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
