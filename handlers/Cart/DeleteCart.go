package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	response "github.com/sabhari/product_catlog/Response"
	services "github.com/sabhari/product_catlog/Services"
	shorthandhelpers "github.com/sabhari/product_catlog/ShorthandHelpers"
	"github.com/sabhari/product_catlog/helpers"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")
	product := urlQuery.Get("product_id")

	product_id, err := strconv.Atoi(product)
	if err != nil {
		helpers.HandleError(response.AtoiErr, err)
		json.NewEncoder(w).Encode(response.AtoiErr)
	}

	found, _ := shorthandhelpers.GetCartHelper(reference, product_id)

	if found == 404 {
		json.NewEncoder(w).Encode(response.ProductDelErr)
		helpers.HandleError(response.ProductDelErr["response"], nil)
		return
	}

	status := services.DeleteCartBL(reference, product_id)

	if status == 200 {
		//helpers.ResponseWriteToScreen(err, "DELETE ON CART", w)
		json.NewEncoder(w).Encode(response.DeleteCart)
		helpers.HandleError(response.DeleteCart["response"], nil)
	} else if status == 404 {
		json.NewEncoder(w).Encode(response.DeleteCartErr)
		helpers.HandleError(response.DeleteCartErr, nil)
	}

}
