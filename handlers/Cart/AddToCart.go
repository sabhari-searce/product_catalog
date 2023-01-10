package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	query "github.com/sabhari/product_catlog/Queries"
	rsp "github.com/sabhari/product_catlog/Response"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	reference := urlQuery.Get("ref")
	pro := urlQuery.Get("product")
	quan := urlQuery.Get("quantity")

	if reference == "" || pro == "" || quan == "" {
		json.NewEncoder(w).Encode(rsp.CartInEmptyErr)
		helpers.HandleError(rsp.CartInEmptyErr, nil)
		return
	}

	product, err := strconv.Atoi(pro)
	helpers.HandleError(rsp.AtoiErr, err)
	json.NewEncoder(w).Encode(rsp.AtoiErr)

	quantity, err := strconv.Atoi(quan)
	helpers.HandleError(rsp.AtoiErr, err)
	json.NewEncoder(w).Encode(rsp.AtoiErr)

	if product < 0 || quantity < 0 {
		helpers.HandleError(rsp.CartInInvalid, err)
		json.NewEncoder(w).Encode(rsp.CartInInvalid)

		return
	}

	row, err := helpers.RunQuery(query.GetReference, reference)
	helpers.HandleError(rsp.ReferenceGetErr, err)
	json.NewEncoder(w).Encode(rsp.ReferenceGetErr)
	row.Scan()
	if !row.Next() {
		json.NewEncoder(w).Encode(rsp.ReferenceNotFound)
		helpers.HandleError(rsp.ReferenceNotFound, nil)
		return
	}

	row, err = helpers.RunQuery(query.GetInventory, product)
	helpers.HandleError(rsp.GetProductErr["response"], err)
	json.NewEncoder(w).Encode(rsp.GetProductErr)

	row.Scan()
	if !row.Next() {
		json.NewEncoder(w).Encode(rsp.ProductNotFound)
		helpers.HandleError(rsp.ProductNotFound, nil)
		return
	}

	var inv_quantity int

	row, err = helpers.RunQuery(query.GetQuantityInventory, product)
	helpers.HandleError(rsp.ProductNotFound, err)
	json.NewEncoder(w).Encode(rsp.ProductNotFound)

	if row.Next() {
		row.Scan(&inv_quantity)

	}

	if quantity > inv_quantity {
		res := fmt.Sprintln("THE ENTERED QUANTITY IS NOT AVAILABLE IN THE INVENTORY! ONLY ", inv_quantity, " QUANTITY IS AVAILABLE")
		json.NewEncoder(w).Encode(res)
		return
	}

	rows, err := helpers.RunQuery(query.GetCartItem, reference, product)
	helpers.HandleError(rsp.CartItemGetErr, err)
	json.NewEncoder(w).Encode(rsp.CartItemGetErr)
	rows.Scan()

	if rows.Next() {
		new_cart := typedefs.Cart{}
		err := rows.Scan(&new_cart.Reference, &new_cart.ProductID, &new_cart.Quantity)
		helpers.HandleError(rsp.GetRowErr, err)
		json.NewEncoder(w).Encode(rsp.GetRowErr)

		_, err = helpers.RunQuery(query.UpdateCartQuantity, quantity+new_cart.Quantity, reference, product)
		if err != nil {
			helpers.HandleError(rsp.CartItemGetErr, err)
			json.NewEncoder(w).Encode(rsp.CartItemGetErr)
		}
		helpers.ResponseWriteToScreen(err, "THE PRODUCT ALREADY FOUND AND THE QUANTITY HAD BEEN UPDATED FOR IT", w)

	} else {

		_, err = helpers.RunQuery(query.CartInsert, reference, product, quantity)
		if err != nil {
			helpers.HandleError(rsp.CartItemGetErr, err)
			json.NewEncoder(w).Encode(rsp.CartItemGetErr)
		}
		helpers.ResponseWriteToScreen(err, "THE PRODUCT HAS BEEN ADDED TO CART!", w)
	}

	new_quantity := inv_quantity - quantity

	_, err = helpers.RunQuery(query.UpdateCartQuantity, new_quantity, product)
	if err != nil {
		helpers.HandleError(rsp.UpdateInvErr, err)
		json.NewEncoder(w).Encode(rsp.UpdateInvErr)
	}

	if new_quantity == 0 {

		_, err = helpers.RunQuery(query.DeleteInventory, product)
		if err != nil {
			helpers.HandleError(rsp.UpdateInvErr, err)
			json.NewEncoder(w).Encode(rsp.UpdateInvErr)
		}
	}

}
