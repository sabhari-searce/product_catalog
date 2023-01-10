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

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	helpers.HandleError("Conversion Error", err)

	if id <= 0 {

		json.NewEncoder(w).Encode(rsp.ProductIdErr)
		helpers.HandleError(rsp.ProductIdErr, nil)
	}

	found := shorthandhelpers.GetInventoryHelper(id)

	if found == 404 {
		json.NewEncoder(w).Encode(rsp.ProductDelErr)
		helpers.HandleError(rsp.ProductDelErr["response"], nil)
		return
	}

	status := service.DeleteInventoryBL(args["id"])

	if status == 200 {
		json.NewEncoder(w).Encode(rsp.InventoryDel)
		helpers.HandleError(rsp.InventoryDel["response"], nil)

	} else if status == 404 {
		json.NewEncoder(w).Encode(rsp.InventoryDelErr)
		helpers.HandleError(rsp.InventoryDelErr["response"], nil)
	}

}
