package dao

import (
	"context"
	"currency-conversion-service/util"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"strconv"
	"sync"
)

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

var (
	DynamoClient *dynamodb.Client
	once         sync.Once
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func (d *DynamoDBClient) UpdateRateInDB(rate ExchangeRate) error {
	return UpdateRateInDB(rate)
}

func ConnectDB() (*dynamodb.Client, error) {
	var err error
	once.Do(func() {
		cfg, loadErr := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(util.AppConfig.AWS.Region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(util.AppConfig.AWS.AccessKeyID, util.AppConfig.AWS.SecretAccessKey, "")),
			config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: util.AppConfig.Server.Endpoint}, nil
			})),
		)

		if loadErr != nil {
			log.Printf("Failed to load AWS config: %v", loadErr)
			err = loadErr
			return
		}

		DynamoClient = dynamodb.NewFromConfig(cfg)
		log.Println("DynamoDB client successfully initialized")
	})

	if DynamoClient == nil {
		return nil, fmt.Errorf("DynamoDB client initialization failed")
	}

	return DynamoClient, err
}

func UpdateRateInDB(rate ExchangeRate) error {
	if DynamoClient == nil {
		return fmt.Errorf("DynamoClient is not initialized")
	}

	_, err := DynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("conversion_rates"),
		Item: map[string]types.AttributeValue{
			"currency": &types.AttributeValueMemberS{Value: rate.Currency},
			"rate":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", rate.Rate)},
		},
	})
	return err
}

func GetRate(currency string) (float64, error) {
	if DynamoClient == nil {
		return 0, fmt.Errorf("DynamoClient is not initialized")
	}

	result, err := DynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("conversion_rates"),
		Key: map[string]types.AttributeValue{
			"currency": &types.AttributeValueMemberS{Value: currency},
		},
	})
	if err != nil {
		return 0, err
	}

	rate, err := strconv.ParseFloat(result.Item["rate"].(*types.AttributeValueMemberN).Value, 64)
	if err != nil {
		return 0, err
	}
	return rate, nil
}
