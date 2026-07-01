package main

import (
	"fmt"
	//"io"
	"net/http"
)

type Config struct {
	URLs string `json:"urls"`
	//Interval int    `json:"interval"`
	//Token    string `json:"token"`
	//ChatId   string `json:"chat_id"`
}

var ConfigList = make([]Config, 0)

var urls = []string{
	"https://www.google.com",
	"https://www.facebook.com",
	"https://www.twitter.com",
	"https://www.instagram.com",
	"https://www.reddit.com",
	"https://www.onliner.by",
}

func GetConfig(arr *[]string) (*[]Config, error) {
	for _, url := range *arr {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		c := Config{
			URLs: resp.Status,
			//Interval: 0,
			//Token:    "",
			//ChatId:   "",
		}
		ConfigList = append(ConfigList, c)
	}
	return &ConfigList, nil
}
