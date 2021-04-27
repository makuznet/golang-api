package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Hello, World!")

	dbinfo := fmt.Sprintf("host=10.0.2.15 user=api password=netlab dbname=api sslmode=disable")

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT name from roles")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", name)
	}

	defer db.Close()
}
