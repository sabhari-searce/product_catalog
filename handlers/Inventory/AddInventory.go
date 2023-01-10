package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gookit/validate"
	rsp "github.com/sabhari/product_catlog/Response"
	service "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	inventory := typedefs.Inventory{}
	json.NewDecoder(r.Body).Decode(&inventory)

	v := validate.Struct(inventory)
	// v := validate.New(u)
	if v.Validate() { // validate ok
		if inventory.ProductID <= 0 {
			json.NewEncoder(w).Encode(rsp.ProductIdErr)
			helpers.HandleError(rsp.ProductIdErr, nil)
			return
		}

		if inventory.Quantity < 0 {
			json.NewEncoder(w).Encode(rsp.QuantityNeg)
			helpers.HandleError(rsp.QuantityNeg, nil)
			return
		}

		status := service.AddInventoryBL(inventory)

		if status == 200 {
			json.NewEncoder(w).Encode(rsp.InventoryIn)
			helpers.HandleError(rsp.InventoryIn, nil)
		} else if status == 404 {
			json.NewEncoder(w).Encode(rsp.InventoryInErr)
			helpers.HandleError(rsp.InventoryInErr, nil)
		}
	} else {
		fmt.Println(v.Errors) // all error messages
	}

}
