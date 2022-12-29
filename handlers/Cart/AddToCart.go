package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	reference := urlQuery.Get("ref")
	pro := urlQuery.Get("product")
	quan := urlQuery.Get("quantity")

	if reference == "" || pro == "" || quan == "" {
		json.NewEncoder(w).Encode("EMPTY QUERY FOUND FOR INSERTING IN CART")
		return
	}

	product, err := strconv.Atoi(pro)
	helpers.HandleError("Error in conversion", err, w)

	quantity, err := strconv.Atoi(quan)
	helpers.HandleError("Error in conversion", err, w)

	if product < 0 || quantity < 0 {
		json.NewEncoder(w).Encode("NEGATIVE VALUE FOUND!!")
		return
	}

	query := "SELECT * FROM cart_reference WHERE ref=$1"

	row, err := helpers.RunQuery(query, w, reference)
	helpers.HandleError("Error in getting reference", err, w)
	row.Scan()
	if !row.Next() {
		json.NewEncoder(w).Encode("THE ENTERED REFERENCE ID NOT CREATED")
		return
	}

	query = "SELECT * FROM inventory WHERE product_id=$1"

	row, err = helpers.RunQuery(query, w, product)
	helpers.HandleError("Error in getting reference", err, w)
	row.Scan()
	if !row.Next() {
		json.NewEncoder(w).Encode("THE ENTERED PRODUCT NOT FOUND")
		return
	}

	var inv_quantity int

	query = "SELECT quantity FROM inventory WHERE product_id=$1"

	row, err = helpers.RunQuery(query, w, product)
	helpers.HandleError("Error in getting reference", err, w)

	if row.Next() {
		row.Scan(&inv_quantity)

	}

	if quantity > inv_quantity {
		res := fmt.Sprintln("THE ENTERED QUANTITY IS NOT AVAILABLE IN THE INVENTORY! ONLY ", inv_quantity, " QUANTITY IS AVAILABLE")
		json.NewEncoder(w).Encode(res)
		return
	}

	query = "SELECT * FROM cart_item WHERE ref=$1 AND product_id=$2"

	rows, err := helpers.RunQuery(query, w, reference, product)
	helpers.HandleError("Error in getting cart item", err, w)
	rows.Scan()

	if rows.Next() {
		new_cart := typedefs.Cart{}
		err := rows.Scan(&new_cart.Reference, &new_cart.ProductID, &new_cart.Quantity)
		helpers.HandleError("Error in rows next", err, w)

		query = "UPDATE cart_item SET quantity=$1 WHERE ref=$2 AND product_id=$3"
		_, err = helpers.RunQuery(query, w, quantity+new_cart.Quantity, reference, product)
		helpers.HandleError("Error in getting cart item", err, w)
		helpers.ResponseWriteToScreen(err, "THE PRODUCT ALREADY FOUND AND THE QUANTITY HAD BEEN UPDATED FOR IT", w)

	} else {
		query = "INSERT INTO cart_item VALUES($1,$2,$3)"
		_, err = helpers.RunQuery(query, w, reference, product, quantity)
		helpers.HandleError("Error in getting cart item", err, w)
		helpers.ResponseWriteToScreen(err, "THE PRODUCT HAS BEEN ADDED TO CART!", w)
	}

	new_quantity := inv_quantity - quantity

	query = "UPDATE inventory SET quantity=$1 WHERE product_id=$2"
	_, err = helpers.RunQuery(query, w, new_quantity, product)
	helpers.HandleError("ERROR IN UPDATING INVENTORY", err, w)

	if new_quantity == 0 {

		query = "DELETE FROM inventory WHERE product_id=$1"
		_, err = helpers.RunQuery(query, w, product)
		helpers.HandleError("ERROR IN UPDATING INVENTORY", err, w)
	}

}
