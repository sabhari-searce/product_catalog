package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	shorthandhelpers "github.com/sabhari/product_catlog/ShorthandHelpers"
	"github.com/sabhari/product_catlog/helpers"
)

var key_elements []string = []string{"quantity"}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	var cart map[string]any
	json.NewDecoder(r.Body).Decode(&cart)
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")
	product := urlQuery.Get("product_id")
	//fmt.Println("The product is", product, urlQuery)
	product_id, err := strconv.Atoi(product)
	helpers.HandleError("Str to Int conversion", err, w)

	//fmt.Println("Before checking", reference, product_id)
	status, _ := shorthandhelpers.GetCartHelper(reference, product_id, w)
	if status == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED ID NOT FOUND FOR UPDATING"})
		return
	}

	query := "UPDATE cart_item SET "

	for key, value := range cart {

		if !contains(key) {
			json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED KEY NOT FOUND IN THE TABLE"})
			fmt.Println("ENTERED KEY NOT FOUND IN THE TABLE")
			return
		}
		query = query + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE product_id=" + urlQuery.Get("product_id") + " AND ref='" + urlQuery.Get("ref") + "'"
		fmt.Println(query)
		_, erro := helpers.RunQuery(query, w)

		helpers.HandleError("ERROR IN RUNNING UPDATE", erro, w)
		// res_string := fmt.Sprintln(key, " UPDATED ON CART")
		// helpers.ResponseWriteToScreen(erro, res_string, w)
		if erro != nil {
			json.NewEncoder(w).Encode(erro)
			return
		}
		query = "UPDATE cart SET "

	}

	json.NewEncoder(w).Encode(map[string]string{"response": "UPDATED ON CART DONE"})

}

func contains(word string) bool {
	for _, elem := range key_elements {
		if word == elem {
			return true
		}
	}
	return false
}
