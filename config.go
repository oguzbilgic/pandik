package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Monitors []*Monitor `json:"checks"`
}

func parseConfig(path *string) (*Config, error) {
	configFile, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
