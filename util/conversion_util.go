package util

import (
	"encoding/json"
	"io/ioutil"
)

var ConversionRates map[string]float64

func LoadConversionRates(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &ConversionRates)
}
