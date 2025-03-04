package util

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	API struct {
		URL string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"api"`
	AWS struct {
		Region          string `yaml:"region"`
		AccessKeyID     string `yaml:"accessKeyID"`
		SecretAccessKey string `yaml:"secretAccessKey"`
	} `yaml:"aws"`
	Server struct {
		Endpoint string `yaml:"endpoint"`
	} `yaml:"server"`
}

var AppConfig Config

func LoadConfig(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}

	log.Printf("Config loaded: %+v\n", AppConfig)
	return nil
}
