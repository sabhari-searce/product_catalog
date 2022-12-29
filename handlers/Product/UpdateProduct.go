package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	shorthandhelpers "github.com/sabhari/product_catlog/ShorthandHelpers"
	"github.com/sabhari/product_catlog/helpers"
)

var key_elements []string = []string{"name", "specification", "sku", "category_id", "price"}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product map[string]any
	json.NewDecoder(r.Body).Decode(&product)
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	helpers.HandleError("Eror when converting string to int", err, w)

	status, _ := shorthandhelpers.GetProductHelper(id, w)
	if status == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED ID NOT FOUND FOR UPDATING"})
		return
	}

	query := "UPDATE product SET "

	for key, value := range product {

		if key == "specification" {
			new_value, err := json.Marshal(value.(map[string]any))
			query = query + key + "='" + string(new_value) + "'" + " WHERE product_id=" + args["id"]
			helpers.HandleError("ERROR IN JSON CONVERSION", err, w)
			_, erro := helpers.RunQuery(query, w)
			helpers.HandleError("ERROR IN RUNNING UPDATE", erro, w)
			//helpers.ResponseWriteToScreen(erro, "SPECIFICATION UPDATED ON PRODUCT", w)
			// if erro == nil {
			// 	json.NewEncoder(w).Encode(map[string]string{"response": "UPDATED ON PRODUCT DONE"})
			// }
		} else {
			if !contains(key) {
				json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED KEY NOT FOUND IN THE TABLE"})
				fmt.Println("ENTERED KEY NOT FOUND IN THE TABLE")
				return
			}
			query = query + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE product_id=" + args["id"]
			_, erro := helpers.RunQuery(query, w)
			helpers.HandleError("ERROR IN RUNNING UPDATE", erro, w)

			if erro != nil {
				json.NewEncoder(w).Encode(erro)
				return
			}
			query = "UPDATE product SET "

		}

	}

	json.NewEncoder(w).Encode(map[string]string{"response": "UPDATED ON PRODUCT DONE"})

}

func contains(word string) bool {
	for _, elem := range key_elements {
		if word == elem {
			return true
		}
	}
	return false
}
