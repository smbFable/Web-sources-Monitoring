package main

import "fmt"

func main() {
	cfg, err := GetConfig(&urls)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
