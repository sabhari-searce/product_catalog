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

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	var cart map[string]any
	json.NewDecoder(r.Body).Decode(&cart)
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")
	product := urlQuery.Get("product_id")
	if reference == "" {
		helpers.HandleError(response.RefErr, nil)
		json.NewEncoder(w).Encode(response.RefErr)
	}
	//fmt.Println("The product is", product, urlQuery)
	product_id, err := strconv.Atoi(product)
	if err != nil {
		helpers.HandleError(response.AtoiErr, err)
		json.NewEncoder(w).Encode(response.AtoiErr)

	}

	if product_id <= 0 {
		helpers.HandleError(response.ProductIdErr, err)
		json.NewEncoder(w).Encode(response.ProductIdErr)
	}

	//fmt.Println("Before checking", reference, product_id)
	status, _ := shorthandhelpers.GetCartHelper(reference, product_id)
	if status == 404 {
		json.NewEncoder(w).Encode(response.CartIdErr)
		helpers.HandleError(response.CartIdErr, nil)
		return
	}

	status = services.UpdateCartBL(cart, product, reference)

	if status == 200 {
		json.NewEncoder(w).Encode(response.UpdateCart)
		helpers.HandleError(response.UpdateCart["response"], nil)
	} else if status == 404 {
		json.NewEncoder(w).Encode(response.UpdateCartErr)
		helpers.HandleError(response.UpdateCartErr, nil)
	} else if status == 403 {
		json.NewEncoder(w).Encode(response.UpdateKeyCart)
		helpers.HandleError(response.UpdateKeyCart, nil)
	}

}
