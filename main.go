package main

import "fmt"

func main() {
	data, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}
