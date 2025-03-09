package consumer

import (
	"currency-conversion-service/consumer/mocks"
	"currency-conversion-service/dao"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"strings"
	"testing"
	"time"
)

func setupTest(kafkaMessage *kafka.Message, kafkaError, dbError error) (*mocks.MockKafkaConsumer, *mocks.MockDatabase, *strings.Builder) {
	mockConsumer := &mocks.MockKafkaConsumer{
		ReadMessageFunc: func(timeout time.Duration) (*kafka.Message, error) {
			return kafkaMessage, kafkaError
		},
	}

	mockDB := &mocks.MockDatabase{
		UpdateRateInDBFunc: func(rate dao.ExchangeRate) error {
			return dbError
		},
	}

	logOutput := &strings.Builder{}
	log.SetOutput(logOutput)

	return mockConsumer, mockDB, logOutput
}

func runConsumer(consumer KafkaConsumer, db Database) {
	done := make(chan bool)
	go func() {
		ConsumeMessages(consumer, db)
		done <- true
	}()

	select {
	case <-done:
		// This shouldn't happen in normal operation
	case <-time.After(100 * time.Millisecond):
		// This is expected; we're manually stopping after a short time
	}
}

func TestConsumeMessages_ValidMessage(t *testing.T) {
	mockConsumer, mockDB, logOutput := setupTest(
		&kafka.Message{Value: []byte(`{"currency":"USD","rate":1.23}`)},
		nil,
		nil,
	)

	runConsumer(mockConsumer, mockDB)

	expectedOutput := "Updated rate: USD = 1.230000"
	if !strings.Contains(logOutput.String(), expectedOutput) {
		t.Errorf("Expected log output to contain '%s', got '%s'", expectedOutput, logOutput.String())
	}
}

func TestConsumeMessages_KafkaError(t *testing.T) {
	mockConsumer, mockDB, logOutput := setupTest(
		nil,
		kafka.NewError(kafka.ErrTimedOut, "Timed out", false),
		nil,
	)

	runConsumer(mockConsumer, mockDB)

	expectedOutput := "Error reading Kafka message: Timed out"
	if !strings.Contains(logOutput.String(), expectedOutput) {
		t.Errorf("Expected log output to contain '%s', got '%s'", expectedOutput, logOutput.String())
	}
}

func TestConsumeMessages_InvalidJSON(t *testing.T) {
	mockConsumer, mockDB, logOutput := setupTest(
		&kafka.Message{Value: []byte(`invalid json`)},
		nil,
		nil,
	)

	runConsumer(mockConsumer, mockDB)

	expectedOutput := "Failed to parse Kafka message:"
	if !strings.Contains(logOutput.String(), expectedOutput) {
		t.Errorf("Expected log output to contain '%s', got '%s'", expectedOutput, logOutput.String())
	}
}

func TestConsumeMessages_DatabaseError(t *testing.T) {
	mockConsumer, mockDB, logOutput := setupTest(
		&kafka.Message{Value: []byte(`{"currency":"EUR","rate":0.89}`)},
		nil,
		errors.New("database error"),
	)

	runConsumer(mockConsumer, mockDB)

	expectedOutput := "Failed to update database: database error"
	if !strings.Contains(logOutput.String(), expectedOutput) {
		t.Errorf("Expected log output to contain '%s', got '%s'", expectedOutput, logOutput.String())
	}
}

func TestConsumeMessages_MultipleMessages(t *testing.T) {
	messages := []*kafka.Message{
		{Value: []byte(`{"currency":"USD","rate":1.23}`)},
		{Value: []byte(`{"currency":"EUR","rate":0.89}`)},
	}
	messageIndex := 0

	mockConsumer := &mocks.MockKafkaConsumer{
		ReadMessageFunc: func(timeout time.Duration) (*kafka.Message, error) {
			if messageIndex < len(messages) {
				msg := messages[messageIndex]
				messageIndex++
				return msg, nil
			}
			return nil, kafka.NewError(kafka.ErrTimedOut, "Timed out", false)
		},
	}

	mockDB := &mocks.MockDatabase{
		UpdateRateInDBFunc: func(rate dao.ExchangeRate) error {
			return nil
		},
	}

	logOutput := &strings.Builder{}
	log.SetOutput(logOutput)

	runConsumer(mockConsumer, mockDB)

	expectedOutputs := []string{
		"Updated rate: USD = 1.230000",
		"Updated rate: EUR = 0.890000",
	}

	for _, expectedOutput := range expectedOutputs {
		if !strings.Contains(logOutput.String(), expectedOutput) {
			t.Errorf("Expected log output to contain '%s', got '%s'", expectedOutput, logOutput.String())
		}
	}
}
