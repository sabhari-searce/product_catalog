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

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	helpers.HandleError(rsp.AtoiErr, err)

	if id <= 0 {
		json.NewEncoder(w).Encode(rsp.CategoryIdErr)
		helpers.HandleError(rsp.CategoryIdErr, nil)
	}
	found := shorthandhelpers.GetCategoryHelper(id)

	if found == 404 {
		json.NewEncoder(w).Encode(rsp.CategoryDelNotFound)
		helpers.HandleError(rsp.CategoryDelNotFound["response"], nil)
		return
	}

	status := service.DeleteCategoryBL(args)

	if status == 404 {
		json.NewEncoder(w).Encode(rsp.CategoryDelError)
		helpers.HandleError(rsp.CategoryDelError, nil)

	} else if status == 200 {
		json.NewEncoder(w).Encode(rsp.CategoryDel)
		helpers.HandleError(rsp.CategoryDel["response"], nil)
	}

}
