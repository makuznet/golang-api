package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// A global var used when connecting to a db
var db *sql.DB

type Product struct {
	Id          int
	Title       string
	Price       int
	Description string
	Category    string
	Image       string
}

type Products struct {
	Products []Product
}

func main() {

	var err error

	// This allows to have just one connection to a db
	db, err = sql.Open("postgres", "host=10.0.2.15 user=api password=netlab dbname=api sslmode=disable")
	if err != nil {
		panic(err)
	}

	// do not close a connection to a db until main function is closed
	defer db.Close()

	fmt.Println("Server started")

	http.HandleFunc("/v1/products/add", addProduct)
	http.HandleFunc("/v1/products/", getProducts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	} else {
		fmt.Fprintf(w, "Product was edded successfully\n")
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	} else {
		w_array := Products{}

		fmt.Println("# Querying")
		rows, err := db.Query("SELECT id,title,price,description,category,image from products")
		if err != nil {
			panic(err)
		}

		for rows.Next() {

			w_product := Product{}

			err = rows.Scan(&w_product.Id, &w_product.Title, &w_product.Price, &w_product.Description, &w_product.Category, &w_product.Image)
			if err != nil {
				panic(err)
			}
			w_array.Products = append(w_array.Products, w_product)
		}

		json.NewEncoder(w).Encode(w_array)
	}
}
