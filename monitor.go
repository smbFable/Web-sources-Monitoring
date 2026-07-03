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

func SiteReliability(url string) (PingResult, error) {
	start := time.Now()
	data, err := http.Get(url)
	if err != nil {
		return PingResult{URL: url}, err
	}
	defer data.Body.Close()
	isup := false
	if data.StatusCode == 200 {
		isup = true
	}

	p := PingResult{
		URL:      url,
		IsUp:     isup,
		PingTime: time.Since(start),
		Status:   data.Status,
	}
	return p, nil
}
