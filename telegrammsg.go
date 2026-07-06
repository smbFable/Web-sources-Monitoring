package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func SendMessage(token string, chatID string, message string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		token, chatID, url.QueryEscape(message))

	resp, err := client.Get(apiURL)
	if err != nil {
		fmt.Println("Ошибка отправки алерта в Telegram: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Telegram вернул ошибку, статус: %s\n", resp.Status)
	}
}
