package server

import (
	"context"
	pb "currency-conversion-service/proto/moneyconverter"
	"currency-conversion-service/server/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert_Success(t *testing.T) {
	s := &server{converter: &mocks.MockConverter{}}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "USD",
			Amount:   100,
		},
		ToCurrency: "EUR",
	}

	resp, err := s.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "EUR", resp.Converted.Currency)
	assert.Equal(t, float64(85), resp.Converted.Amount) // Use float64 here
}

func TestConvert_UnsupportedCurrency(t *testing.T) {
	s := &server{converter: &mocks.MockConverter{}}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "USD",
			Amount:   100,
		},
		ToCurrency: "XYZ",
	}

	_, err := s.Convert(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, "unsupported currency", err.Error())
}
