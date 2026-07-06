package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("Ошибка загрузки конфига:", err)
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	fmt.Println("Сервис мониторинга успешно запущен!")
	fmt.Printf("Интервал проверок: %d секунд\n", cfg.Interval)

	ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	defer ticker.Stop()

	RunScan(cfg, httpClient)

	for {
		select {
		case <-ticker.C:
			RunScan(cfg, httpClient)
		}
	}
}

func RunScan(cfg *Config, client *http.Client) {
	fmt.Println("\n Начинаем сканирование сайтов...")

	urlchan := make(chan PingResult, len(cfg.URLs))
	var arrchan []PingResult

	for _, u := range cfg.URLs {
		go SiteReliability(client, u, urlchan)
	}

	for i := 0; i < len(cfg.URLs); i++ {
		res := <-urlchan
		arrchan = append(arrchan, res)

		if !res.IsUp {
			alertMsg := fmt.Sprintf("ВНИМАНИЕ! Сайт упал!\nURL: %s\nСтатус: %s\nПинг: %v",
				res.URL, res.Status, res.PingTime)
			go SendMessage(cfg.Token, cfg.ChatId, alertMsg)
		}
	}

	for i, res := range arrchan {
		statusIcon := "🟢"
		if !res.IsUp {
			statusIcon = "🔴"
		}
		fmt.Printf("%d. %s %s - %s (%v)\n", i+1, statusIcon, res.URL, res.Status, res.PingTime)
	}
	fmt.Println("Сканирование завершено. Ожидание...")
}
