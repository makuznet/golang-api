package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	http.HandleFunc("/v1/products/", getProduct)
	http.HandleFunc("/v1/products/add", addProduct)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Get product info: all or by a product id
func getProduct(w http.ResponseWriter, r *http.Request) {

	productId, _ := strconv.Atoi(r.URL.Path[13:])

	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	} else if productId != 0 {

		fmt.Println("Getting product info by userId:", productId)

		rows, err := db.Query("SELECT id, title, price, description, category, image from products where id=$1", productId)
		if err != nil {
			panic(err)
		}

		w_array := Products{}

		for rows.Next() {

			w_product := Product{}

			err = rows.Scan(&w_product.Id, &w_product.Title, &w_product.Price, &w_product.Description, &w_product.Category, &w_product.Image)
			if err != nil {
				panic(err)
			}
			w_array.Products = append(w_array.Products, w_product)
		}

		json.NewEncoder(w).Encode(w_array)

	} else {

		fmt.Println("Getting the list of products")
		rows, err := db.Query("SELECT id,title,price,description,category,image from products")
		if err != nil {
			panic(err)
		}

		w_array := Products{}

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

// Add new product
func addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	} else {

		decoder := json.NewDecoder(r.Body)

		var g_product Product

		err := decoder.Decode(&g_product)
		if err != nil {
			panic(err)
		}

		query := fmt.Sprintf("INSERT INTO products(title, price, description, category, image) VALUES('%s', %d, '%s', '%s', '%s') RETURNING id", g_product.Title, g_product.Price, g_product.Description, g_product.Category, g_product.Image)

		fmt.Println("# Inserting")
		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}

		for rows.Next() {

			var id int

			err = rows.Scan(&id)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "{\"id\":%d}", id)
		}
	}
}
