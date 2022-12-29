package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	cart "github.com/sabhari/product_catlog/handlers/Cart"
	category "github.com/sabhari/product_catlog/handlers/Category"
	inventory "github.com/sabhari/product_catlog/handlers/Inventory"
	product "github.com/sabhari/product_catlog/handlers/Product"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/product/add", product.AddProduct).Methods("POST")
	router.HandleFunc("/product/get/{id:[0-9]+}", product.GetProduct).Methods("GET")
	router.HandleFunc("/product/delete/{id:[0-9]+}", product.DeleteProduct).Methods("GET")
	router.HandleFunc("/product/getpage/{id:[0-9]+}", product.GetShortProduct).Methods("GET")
	router.HandleFunc("/product/update/{id:[0-9]+}", product.UpdateProduct).Methods("POST")

	router.HandleFunc("/category/add", category.AddCategory).Methods("POST")
	router.HandleFunc("/category/get/{id:[0-9]+}", category.GetCategory).Methods("GET")
	router.HandleFunc("/category/delete/{id:[0-9]+}", category.DeleteCategory).Methods("GET")
	router.HandleFunc("/category/update/{id:[0-9]+}", category.UpdateCategory).Methods("POST")

	router.HandleFunc("/inventory/add", inventory.AddInventory).Methods("POST")
	router.HandleFunc("/inventory/get/{id:[0-9]+}", inventory.GetInventory).Methods("GET")
	router.HandleFunc("/inventory/delete/{id:[0-9]+}", inventory.DeleteInventory).Methods("GET")
	router.HandleFunc("/inventory/update/{id:[0-9]+}", inventory.UpdateInventory).Methods("POST")

	router.HandleFunc("/cart/create", cart.CreateCartReference).Methods("POST")
	router.HandleFunc("/cart/add", cart.AddToCart).Methods("POST")
	router.HandleFunc("/cart/get", cart.GetCart).Methods("GET")
	router.HandleFunc("/cart/delete", cart.DeleteCategory).Methods("GET")
	router.HandleFunc("/cart/update", cart.UpdateCart).Methods("POST")
	router.HandleFunc("/cart/additems", cart.AddItemsToCart).Methods("POST")

	fmt.Println("SERVER STARTED ON PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
