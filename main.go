package main

import (
	"fmt"
	"log"
)

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	pr, err := SiteReliability(cfg.URLs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pr)
}
