package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Money struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type ConvertRequest struct {
	From       Money  `json:"from"`
	ToCurrency string `json:"to_currency"`
}

type ConvertResponse struct {
	Converted Money `json:"converted"`
}

func main() {
	url := "http://localhost:8080/v1/convert"
	reqBody := &ConvertRequest{
		From: Money{
			Currency: "EUR",
			Amount:   100.0,
		},
		ToCurrency: "USD",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshalling request: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	var res ConvertResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
	}

	fmt.Printf("%.2f %s is %.2f %s\n", reqBody.From.Amount, reqBody.From.Currency, res.Converted.Amount, res.Converted.Currency)
}
