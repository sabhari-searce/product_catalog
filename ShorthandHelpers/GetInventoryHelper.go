package shorthandhelpers

import (
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetInventoryHelper(id int, w http.ResponseWriter) (int, []typedefs.Inventory) {
	query := "SELECT * FROM inventory WHERE product_id=$1"
	rows, err := helpers.RunQuery(query, w, id)
	helpers.HandleError("Error in getting category", err, w)
	rows.Scan()
	list_of_inventory := []typedefs.Inventory{}

	for rows.Next() {
		new_inventory := typedefs.Inventory{}
		err := rows.Scan(&new_inventory.ProductID, &new_inventory.Quantity)
		helpers.HandleError("Error in rows next", err, w)
		list_of_inventory = append(list_of_inventory, new_inventory)
	}

	if len(list_of_inventory) == 0 {
		return 404, list_of_inventory
	}
	return 200, list_of_inventory

}
