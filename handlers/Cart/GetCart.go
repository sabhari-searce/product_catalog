package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")

	query := "SELECT cart_item.ref,cart_item.product_id,product.name,product.specification,category.name,product.price,cart_item.quantity FROM (cart_item JOIN product ON cart_item.product_id = product.product_id) JOIN category ON category.category_id = product.product_id WHERE cart_item.ref=$1"

	list_of_cart := []typedefs.GetCartStruct{}

	rows, err := helpers.RunQuery(query, w, reference)
	rows.Scan()
	helpers.HandleError("Error in getting Category", err, w)
	var total float32

	for rows.Next() {
		new_cart := typedefs.GetCartStruct{}
		spec_json := ""
		err := rows.Scan(&new_cart.Reference, &new_cart.ProductID, &new_cart.ProductName, &spec_json, &new_cart.CategoryName, &new_cart.Price, &new_cart.Quantity)
		helpers.HandleError("Error in rows next", err, w)
		helpers.HandleError("Error in rows next", err, w)
		json.Unmarshal([]byte(spec_json), &new_cart.Specification)
		total += (new_cart.Price * float32(new_cart.Quantity))
		list_of_cart = append(list_of_cart, new_cart)
	}

	if len(list_of_cart) == 0 {
		json.NewEncoder(w).Encode("NO DATA FOUND!")
		return
	}

	err = json.NewEncoder(w).Encode(list_of_cart)
	res := fmt.Sprintln("The total price of this cart is ", total)
	err = json.NewEncoder(w).Encode(res)

	fmt.Println(list_of_cart)
	fmt.Println(res)

}
