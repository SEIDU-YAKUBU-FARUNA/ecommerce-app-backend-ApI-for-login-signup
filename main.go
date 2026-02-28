package main

import (
	"go-ecommerce/database"
	"go-ecommerce/handlers"
	"log"
	"net/http"
)

func main() {

	database.ConnectDB()

	//Auth route
	http.HandleFunc("/register", handlers.RegisterUser)
	http.HandleFunc("/login", handlers.LoginUser)

	//Products route for customer
	http.HandleFunc("/products", handlers.GetProducts)
	http.HandleFunc("/product", handlers.GetProduct)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
