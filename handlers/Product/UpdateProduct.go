package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	rsp "github.com/sabhari/product_catlog/Response"
	service "github.com/sabhari/product_catlog/Services"
	shorthandhelpers "github.com/sabhari/product_catlog/ShorthandHelpers"
	"github.com/sabhari/product_catlog/helpers"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product map[string]any
	json.NewDecoder(r.Body).Decode(&product)
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	helpers.HandleError(rsp.AtoiErr, err)

	if id <= 0 {
		json.NewEncoder(w).Encode(rsp.ProductIdErr)
		helpers.HandleError(rsp.ProductIdErr, nil)
	}

	status := shorthandhelpers.GetProductHelper(id)
	if status == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": rsp.ProductIdErr})
		helpers.HandleError(rsp.ProductIdErr, nil)
		return
	}

	report := service.UpdateProductBL(product, args["id"])

	if report == 200 {
		json.NewEncoder(w).Encode(rsp.ProductUp)
		helpers.HandleError(rsp.ProductUp["response"], nil)

	} else if report == 404 {
		json.NewEncoder(w).Encode(rsp.ProductUpErr)
		helpers.HandleError(rsp.ProductUpErr, nil)
	} else if report == 403 {
		json.NewEncoder(w).Encode(rsp.ProductKeyErr)
	}

}
