package shorthandhelpers

import (
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCartHelper(ref string, id int) (int, []typedefs.Cart) {
	query := "SELECT * FROM cart_item WHERE ref=$1 AND product_id=$2"
	rows, err := helpers.RunQuery(query, ref, id)
	helpers.HandleError("Error in getting category", err)
	rows.Scan()
	list_of_cart := []typedefs.Cart{}

	for rows.Next() {
		new_cart := typedefs.Cart{}
		err := rows.Scan(&new_cart.Reference, &new_cart.ProductID, &new_cart.Quantity)
		helpers.HandleError("Error in rows next", err)
		list_of_cart = append(list_of_cart, new_cart)
	}

	if len(list_of_cart) == 0 {
		return 404, list_of_cart
	}
	return 200, list_of_cart

}
