package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
		fmt.Printf("%s\n", name)
	}

	http.HandleFunc("/v1/products/", getProducts)
	log.Fatal(http.ListenAndServe(":8080", nil))

	defer db.Close()
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"products\": [{\"id\": 2, \"title\": \"Mens Casual Premium Slim Fit T-Shirts\", \"price\": 22.3, \"description\": \"Slim-fitting style, contrast raglan long sleeve, three-button henley placket, light weight & soft fabric for breathable and comfortable wearing. And Solid stitched shirts with round neck made for durability and a great fit for casual fashion wear and diehard baseball fans. The Henley style round neckline includes a three-button placket.\", \"category\": \"mens clothing\",\"image\": \"https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg\"}]}")
}
