package main

import (
	"net/http"
)

func SendAlert(token string, chat_id string, message string) {
	http.Post("https://api.telegram.org/"+token+"/"+chat_id+"/"+message, "application/json", nil)
}
