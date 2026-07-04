package main

import (
	"database/sql"
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

	for index, url := range cfg.URLs {
		pr, err := SiteReliability(url)
		if err != nil {
			fmt.Println(index+1, ": ", pr)
			continue
		}
		fmt.Println(index+1, ": ", pr)
	}

	database, err := sql.Open("sqlite", "./Web.db")
	if err != nil {
		fmt.Println(err)
	}
	table, err := database.Prepare("CREATE TABLE IF NOT EXISTS WEBdb (URL TEXT, IsUp INTEGER, Ping INTEGER, Status TEXT)")
	if err != nil {
		fmt.Println(err)
	}
	defer database.Close()

	_, err = table.Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Таблица создана!")
}
