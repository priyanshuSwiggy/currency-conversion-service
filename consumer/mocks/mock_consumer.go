package mocks

import (
	"currency-conversion-service/dao"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type MockKafkaConsumer struct {
	ReadMessageFunc func(timeout time.Duration) (*kafka.Message, error)
	CloseFunc       func() error
	SubscribeFunc   func(topics []string, rebalanceCb kafka.RebalanceCb) error
}

func (m *MockKafkaConsumer) ReadMessage(timeout time.Duration) (*kafka.Message, error) {
	return m.ReadMessageFunc(timeout)
}

func (m *MockKafkaConsumer) Close() error {
	return m.CloseFunc()
}

func (m *MockKafkaConsumer) SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error {
	return m.SubscribeFunc(topics, rebalanceCb)
}

type MockDatabase struct {
	UpdateRateInDBFunc func(rate dao.ExchangeRate) error
}

func (m *MockDatabase) UpdateRateInDB(rate dao.ExchangeRate) error {
	return m.UpdateRateInDBFunc(rate)
}
