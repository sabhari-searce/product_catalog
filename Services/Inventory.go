package services

import (
	"fmt"

	query "github.com/sabhari/product_catlog/Queries"
	rsp "github.com/sabhari/product_catlog/Response"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddInventoryBL(inventory typedefs.Inventory) int {
	_, err := helpers.RunQuery(query.AddInventory, inventory.ProductID, inventory.Quantity)
	if err != nil {
		return 404
	}
	return 200
}

func GetInventoryBL(list_of_inventory []typedefs.Inventory, id string) {
	rows, err := helpers.RunQuery(query.GetInventory, id)
	rows.Scan()
	helpers.HandleError(rsp.InventoryGetErr, err)

	for rows.Next() {
		new_inventory := typedefs.Inventory{}
		err := rows.Scan(&new_inventory.ProductID, &new_inventory.Quantity)
		helpers.HandleError(rsp.GetRowErr, err)
		list_of_inventory = append(list_of_inventory, new_inventory)
	}
}

func DeleteInventoryBL(id string) int {
	_, err := helpers.RunQuery(query.DeleteInventory, id)
	if err != nil {
		return 404
	}
	return 200

}

var key_elements_inv []string = []string{"quantity"}

func containsInv(word string) bool {
	for _, elem := range key_elements_inv {
		if word == elem {
			return true
		}
	}
	return false
}

func UpdateInventoryBL(inventory map[string]any, id int) int {
	queryexe := query.UpdateInventory

	for key, value := range inventory {

		if !containsInv(key) {
			return 403
		}
		queryexe = queryexe + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE product_id=" + string(id)
		_, erro := helpers.RunQuery(queryexe)
		helpers.HandleError(rsp.UpdateInvErr, erro)

		if erro != nil {
			return 404
		}
		queryexe = query.UpdateInventory

	}
	return 200
}
