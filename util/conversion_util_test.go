package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConversionRates(t *testing.T) {
	err := LoadConversionRates("../conversion_rates.json")
	assert.NoError(t, err)
	assert.NotEmpty(t, ConversionRates)
	assert.Equal(t, 83.0, ConversionRates["USD"])
}

func TestLoadConversionRatesInvalidFile(t *testing.T) {
	err := LoadConversionRates("../invalid_file.json")
	assert.Error(t, err)
}

func TestLoadConversionRatesMalformedJSON(t *testing.T) {
	err := LoadConversionRates("../malformed_conversion_rates.json")
	assert.Error(t, err)
}
