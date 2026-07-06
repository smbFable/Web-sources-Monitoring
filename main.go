package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	urlchan := make(chan PingResult, len(cfg.URLs))
	arrchan := make([]PingResult, 0)

	for _, url := range cfg.URLs {
		go SiteReliability(url, urlchan)
	}

	for i := 0; i < len(cfg.URLs); i++ {
		arrchan = append(arrchan, <-urlchan)
	}

	for i := 0; i < len(cfg.URLs); i++ {
		fmt.Println(i+1, ": ", arrchan[i])
	}
}
