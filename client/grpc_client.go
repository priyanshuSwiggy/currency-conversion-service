package main

import (
	"context"
	pb "currency-conversion-service/proto/moneyconverter"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMoneyConverterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		From: &pb.Money{
			Currency: "EUR",
			Amount:   100.0,
		},
		ToCurrency: "INR",
	}

	res, err := client.Convert(ctx, req)
	if err != nil {
		log.Fatalf("Could not convert: %v", err)
	}
	fmt.Printf("%.2f %s is %.2f %s\n", req.GetFrom().GetAmount(), req.GetFrom().GetCurrency(), res.GetConverted().GetAmount(), res.GetConverted().GetCurrency())

}
