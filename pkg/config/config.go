package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Token string `json:"token"`
}

func ReadConfig(filepath string) (Config, error) {
	var config Config

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, err
	}

	return config, err
}
