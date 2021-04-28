package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Server started")

	http.HandleFunc("/v1/products/", getProducts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("postgres", "host=10.0.2.15 user=api password=netlab dbname=api sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT id, title, price, description, category, image from products")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var title string
		var price int
		var description string
		var category string
		var image string

		err = rows.Scan(&id, &title, &price, &description, &category, &image)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%d\n %s\n %d\n %s\n %s\n %s\n", id, title, price, description, category, image)
	}

	db.Close()
}
