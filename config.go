package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	URLs     []string `json:"urls"`
	Interval int      `json:"interval"`
	Token    string   `json:"token"`
	ChatId   string   `json:"chat_id"`
}

func LoadConfig(name string) (*Config, error) {
	cfg := &Config{}

	jsonfiledata, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonfiledata, cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cfg, nil
}
