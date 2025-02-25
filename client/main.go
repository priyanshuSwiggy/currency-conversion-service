package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "currency-conversion-service/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCurrencyConverterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		FromCurrency: "USD",
		ToCurrency:   "INR",
		Amount:       100.0,
	}

	res, err := client.Convert(ctx, req)
	if err != nil {
		log.Fatalf("Could not convert: %v", err)
	}

	fmt.Printf("%.2f %s is %.2f %s\n", req.Amount, req.FromCurrency, res.ConvertedAmount, req.ToCurrency)
}
