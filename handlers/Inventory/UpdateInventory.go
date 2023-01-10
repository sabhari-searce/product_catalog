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

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory map[string]any
	json.NewDecoder(r.Body).Decode(&inventory)
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	helpers.HandleError(rsp.AtoiErr, err)

	status := shorthandhelpers.GetInventoryHelper(id)
	if status == 404 {
		json.NewEncoder(w).Encode(rsp.UpdateInvIdErr)
		helpers.HandleError(rsp.UpdateInvIdErr["response"], nil)
		return
	}

	status = service.UpdateInventoryBL(inventory, id)

	if status == 200 {
		json.NewEncoder(w).Encode(rsp.UpdateInventoryDone)
		helpers.HandleError(rsp.UpdateInventoryDone["response"], nil)

	} else if status == 404 {
		json.NewEncoder(w).Encode(rsp.UpdateInvErr)
		helpers.HandleError(rsp.UpdateInvErr, nil)
	} else if status == 403 {
		json.NewEncoder(w).Encode(rsp.UpdateInvKey)
		helpers.HandleError(rsp.UpdateInvKey, nil)
	}

}
