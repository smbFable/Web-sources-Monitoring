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

	for index, url := range cfg.URLs {
		pr, err := SiteReliability(url)
		if err != nil {
			fmt.Println(index+1, ": ", pr)
			continue
		}
		fmt.Println(index+1, ": ", pr)
	}

}
