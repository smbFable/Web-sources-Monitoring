package main

import (
	"net/http"
	"time"
)

type PingResult struct {
	URL      string
	IsUp     bool
	PingTime time.Duration
	Status   string
}

func SiteReliability(url string, PR chan PingResult) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	start := time.Now()
	data, err := client.Get(url)
	if err != nil {
		PR <- PingResult{
			URL:      url,
			IsUp:     false,
			PingTime: time.Since(start),
			Status:   "TIMEOUT/ERROR",
		}
		return
	}
	defer data.Body.Close()

	p := PingResult{
		URL:      url,
		IsUp:     data.StatusCode == 200,
		PingTime: time.Since(start),
		Status:   data.Status,
	}
	PR <- p
}
