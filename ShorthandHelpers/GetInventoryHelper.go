package shorthandhelpers

import (
	services "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetInventoryHelper(id int) int {

	list_of_inventory := []typedefs.Inventory{}

	services.GetInventoryBL(list_of_inventory, string(id))

	if len(list_of_inventory) == 0 {
		return 404
	}
	return 200

}
