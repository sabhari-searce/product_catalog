package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	inventory := typedefs.Inventory{}
	json.NewDecoder(r.Body).Decode(&inventory)

	query := "INSERT INTO inventory VALUES($1,$2)"

	if inventory.Quantity < 0 {
		json.NewEncoder(w).Encode("Quantity cannot be negative")
		return
	}

	_, err := helpers.RunQuery(query, w, inventory.ProductID, inventory.Quantity)
	helpers.HandleError("Error in inserting", err, w)

	helpers.ResponseWriteToScreen(err, "Insert to Inventory done", w)
}
