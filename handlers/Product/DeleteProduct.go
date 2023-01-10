package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	rsp "github.com/sabhari/product_catlog/Response"
	services "github.com/sabhari/product_catlog/Services"
	shorthandhelpers "github.com/sabhari/product_catlog/ShorthandHelpers"
	"github.com/sabhari/product_catlog/helpers"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	helpers.HandleError(rsp.AtoiErr, err)

	if id <= 0 {
		helpers.HandleError(rsp.ProductIdErr, nil)
		json.NewEncoder(w).Encode(rsp.ProductIdErr)

	}

	found := shorthandhelpers.GetProductHelper(id)

	if found == 404 {
		json.NewEncoder(w).Encode(rsp.ProductDelErr)
		helpers.HandleError(rsp.ProductDelErr["response"], nil)
		return
	}
	services.DeleteProductBL(id)

	//helpers.ResponseWriteToScreen(err, "DELETE ON PRODUCT", w)
	json.NewEncoder(w).Encode(rsp.ProductDel)
	helpers.HandleError(rsp.ProductDel["response"], nil)

}
