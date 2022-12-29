package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	category := typedefs.Category{}
	json.NewDecoder(r.Body).Decode(&category)

	query := "INSERT INTO category VALUES($1,$2)"

	_, err := helpers.RunQuery(query, w, category.CategoryID, category.Name)
	helpers.HandleError("Error in inserting", err, w)

	helpers.ResponseWriteToScreen(err, "Insert to Category done", w)
}
