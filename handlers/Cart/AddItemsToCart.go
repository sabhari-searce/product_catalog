package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
)

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	response := []map[string]any{}
	request_body := []map[string]int{}

	ref := r.URL.Query().Get("ref")

	err := json.NewDecoder(r.Body).Decode(&request_body)
	helpers.HandleError("ERROR IN DECODING", err, w)

	for _, v := range request_body {
		new_response_item := map[string]any{}
		product_id := v["product_id"]
		quantity := v["quantity"]

		url := "http://localhost:8080/cart/add?ref=" + ref + "&product=" + fmt.Sprint(product_id) + "&quantity=" + fmt.Sprint(quantity)
		//fmt.Println(url)
		_, err = http.Post(url, "application/json", nil)
		helpers.HandleError("ERROR IN POST REQUEST", err, w)

		new_response_item["product_id"] = product_id
		new_response_item["quantity"] = quantity
		new_response_item["response"] = "INSERTED SUCCESFULLY"

		// new_response_item["message"] = queryhelpers.AddToCart(ref, fmt.Sprint(quantity), fmt.Sprint(product_id))["message"]
		response = append(response, new_response_item)
	}

	//helpers.SendResponse(response, w)
	//fmt.Println(response)
	//fmt.Println("GOT INSERTED SUCCESFULLY")

	json.NewEncoder(w).Encode(response)
}
