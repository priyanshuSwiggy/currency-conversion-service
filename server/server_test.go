package main

import (
	"context"
	pb "currency-conversion-service/pb"
	"currency-conversion-service/server/mocks"
	"currency-conversion-service/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertUSDToINR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"USD": 83.0,
	}

	server := &mocks.MockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "USD",
			Amount:   100.0,
		},
		ToCurrency: "INR",
	}

	res, err := server.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 8300.0, res.Converted.Amount)
}

func TestConvertINRToUSD(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"USD": 83.0,
	}

	server := &mocks.MockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "INR",
			Amount:   8300.0,
		},
		ToCurrency: "USD",
	}

	res, err := server.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 100.0, res.Converted.Amount)
}

func TestConvertEURToINR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"EUR": 90.0,
	}

	server := &mocks.MockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "EUR",
			Amount:   90.0,
		},
		ToCurrency: "INR",
	}

	res, err := server.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 8100.0, res.Converted.Amount)
}

func TestConvertINRToEUR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"EUR": 90.0,
	}

	server := &mocks.MockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "INR",
			Amount:   8100.0,
		},
		ToCurrency: "EUR",
	}

	res, err := server.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 90.0, res.Converted.Amount)
}

func TestConvertUSDToEUR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"USD": 83.0,
		"EUR": 90.0,
	}

	server := &mocks.MockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "USD",
			Amount:   100.0,
		},
		ToCurrency: "EUR",
	}

	res, err := server.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.InEpsilon(t, 92.22, res.Converted.Amount, 0.01)
}

func TestConvertEURToUSD(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"USD": 83.0,
		"EUR": 90.0,
	}

	server := &mocks.MockServer{}
	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "EUR",
			Amount:   90.0,
		},
		ToCurrency: "USD",
	}

	res, err := server.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.InEpsilon(t, 97.59, res.Converted.Amount, 0.01)
}
