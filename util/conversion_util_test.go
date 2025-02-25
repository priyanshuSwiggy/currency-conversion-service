package util

import (
	"testing"
)

func TestLoadConversionRates(t *testing.T) {
	err := LoadConversionRates("../conversion_rates.json")
	if err != nil {
		t.Fatalf("Failed to load conversion rates: %v", err)
	}

	if len(ConversionRates) == 0 {
		t.Fatalf("Conversion rates should not be empty")
	}

	if ConversionRates["USD"] != 83.0 {
		t.Errorf("Expected USD rate to be 83.0, got %v", ConversionRates["USD"])
	}
}
