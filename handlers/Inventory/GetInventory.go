package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM inventory WHERE product_id = $1"
	args := mux.Vars(r)
	list_of_inventory := []typedefs.Inventory{}

	rows, err := helpers.RunQuery(query, w, args["id"])
	rows.Scan()
	helpers.HandleError("Error in getting Category", err, w)

	for rows.Next() {
		new_inventory := typedefs.Inventory{}
		err := rows.Scan(&new_inventory.ProductID, &new_inventory.Quantity)
		helpers.HandleError("Error in rows next", err, w)
		list_of_inventory = append(list_of_inventory, new_inventory)
	}

	if len(list_of_inventory) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"response": "NO DATA FOUND!"})
		return
	}

	err = json.NewEncoder(w).Encode(list_of_inventory[0])
	helpers.HandleError("Error in writing inventory to screen", err, w)
	fmt.Println(list_of_inventory[0])

}
