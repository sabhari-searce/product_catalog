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

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category map[string]any
	json.NewDecoder(r.Body).Decode(&category)
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	helpers.HandleError(rsp.AtoiErr, err)

	status := shorthandhelpers.GetCategoryHelper(id)
	if status == 404 {
		json.NewEncoder(w).Encode(rsp.CategoryIdUpdateErr)
		helpers.HandleError(rsp.CategoryIdUpdateErr["response"], nil)
		return
	}

	status = service.UpdateCategoryBL(category, args)
	if status == 200 {
		json.NewEncoder(w).Encode(rsp.CategoryUpdate)
		helpers.HandleError(rsp.CategoryUpdate["response"], nil)
	} else if status == 404 {
		json.NewEncoder(w).Encode(rsp.CategoryUpdateErr)
		helpers.HandleError(rsp.CategoryUpdateErr, nil)
	} else if status == 403 {
		json.NewEncoder(w).Encode(rsp.CategoryUpdateKeyErr)
		helpers.HandleError(rsp.CategoryUpdateKeyErr, nil)
	}

}
