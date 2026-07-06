package main

/*
import (
	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

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
)
*/
