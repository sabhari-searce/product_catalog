package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	product := typedefs.Product{}
	json.NewDecoder(r.Body).Decode(&product)

	query := "INSERT INTO product VALUES($1,$2,$3,$4,$5,$6)"

	spec, err := json.Marshal(product.Specification)
	//fmt.Println("Entered into add product", spec)
	helpers.HandleError("Marshal error", err, w)

	_, err = helpers.RunQuery(query, w, product.Product_ID, product.Name, spec, product.SKU, product.CategoryID, product.Price)
	helpers.HandleError("Error in inserting", err, w)

	helpers.ResponseWriteToScreen(err, "Insert to Product done", w)
}
