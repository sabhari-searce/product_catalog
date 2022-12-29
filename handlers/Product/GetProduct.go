package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	query := "SELECT product.product_id,product.name,product.specification,product.sku,category.name,product.price FROM product JOIN category ON product.category_id = category.category_id WHERE product.product_id = $1"
	args := mux.Vars(r)
	list_of_product := []typedefs.ProductDesc{}

	rows, err := helpers.RunQuery(query, w, args["id"])
	rows.Scan()
	helpers.HandleError("Error in getting product", err, w)

	for rows.Next() {
		new_product := typedefs.ProductDesc{}
		spec_json := ""
		err := rows.Scan(&new_product.Product_ID, &new_product.Name, &spec_json, &new_product.SKU, &new_product.Category_name, &new_product.Price)
		helpers.HandleError("Error in rows next", err, w)
		json.Unmarshal([]byte(spec_json), &new_product.Specification)
		list_of_product = append(list_of_product, new_product)
	}

	if len(list_of_product) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"response": "NO DATA FOUND!"})
		return
	}

	err = json.NewEncoder(w).Encode(list_of_product[0])
	helpers.HandleError("Error in returning product", err, w)
	fmt.Println(list_of_product[0])
	//helpers.ResponseWriteToScreen(err, "FETCHING DATA FROM PRODUCT", w)

}
