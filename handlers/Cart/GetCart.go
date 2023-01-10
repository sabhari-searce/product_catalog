package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	rsp "github.com/sabhari/product_catlog/Response"
	services "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")

	if reference == "" {
		json.NewEncoder(w).Encode(rsp.RefErr)
		helpers.HandleError(rsp.RefErr, nil)

	}

	list_of_cart := []typedefs.GetCartStruct{}

	total := services.GetCartBL(list_of_cart, reference)
	if len(list_of_cart) == 0 {
		json.NewEncoder(w).Encode(rsp.GetProductErr)
		helpers.HandleError(rsp.GetProductErr["response"], nil)
		return
	}

	err := json.NewEncoder(w).Encode(list_of_cart)
	json.NewEncoder(w).Encode(rsp.WritingErr)
	helpers.HandleError(rsp.WritingErr, err)
	res := fmt.Sprintln("The total price of this cart is ", total)
	err = json.NewEncoder(w).Encode(res)

	fmt.Println(list_of_cart)
	fmt.Println(res)

}
