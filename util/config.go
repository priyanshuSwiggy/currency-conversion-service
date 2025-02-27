package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	API struct {
		URL string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"api"`
}

var AppConfig Config

func LoadConfig(configFile string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}
	return nil
}
