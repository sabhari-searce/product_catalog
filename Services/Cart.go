package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	query "github.com/sabhari/product_catlog/Queries"
	rsp "github.com/sabhari/product_catlog/Response"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func CreateCartReferenceBL() uuid.UUID {
	reference := uuid.New()

	_, err := helpers.RunQuery(query.CreateCartReference, reference, time.Now())
	if err != nil {
		return uuid.Nil
	}
	return reference
}

func GetCartBL(list_of_cart []typedefs.GetCartStruct, reference string) float32 {
	rows, err := helpers.RunQuery(query.GetCart, reference)
	rows.Scan()
	helpers.HandleError(rsp.GetCartErr, err)
	var total float32

	for rows.Next() {
		new_cart := typedefs.GetCartStruct{}
		spec_json := ""
		err := rows.Scan(&new_cart.Reference, &new_cart.ProductID, &new_cart.ProductName, &spec_json, &new_cart.CategoryName, &new_cart.Price, &new_cart.Quantity)
		helpers.HandleError(rsp.GetRowErr, err)
		json.Unmarshal([]byte(spec_json), &new_cart.Specification)
		total += (new_cart.Price * float32(new_cart.Quantity))
		list_of_cart = append(list_of_cart, new_cart)
	}
	return total

}

func DeleteCartBL(reference string, product_id int) int {
	_, err := helpers.RunQuery(query.DeleteCart, reference, product_id)
	if err != nil {
		return 404
	}
	return 200

}

var key_elements_cart []string = []string{"quantity"}

func containsCart(word string) bool {
	for _, elem := range key_elements_cart {
		if word == elem {
			return true
		}
	}
	return false
}

func UpdateCartBL(cart map[string]any, product_id string, reference string) int {
	queryexe := query.UpdateCart
	for key, value := range cart {

		if !contains(key) {
			return 403
		}
		queryexe = queryexe + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE product_id=" + product_id + " AND ref='" + reference + "'"
		_, erro := helpers.RunQuery(queryexe)

		if erro != nil {
			return 404
		}
		queryexe = query.UpdateCart

	}
	return 200

}
