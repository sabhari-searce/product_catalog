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

var key_elements []string = []string{"quantity"}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory map[string]any
	json.NewDecoder(r.Body).Decode(&inventory)
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	helpers.HandleError("Eror when converting string to int", err, w)

	status, _ := shorthandhelpers.GetInventoryHelper(id, w)
	if status == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED ID NOT FOUND FOR UPDATING"})
		return
	}

	query := "UPDATE inventory SET "

	for key, value := range inventory {

		if !contains(key) {
			json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED KEY NOT FOUND IN THE TABLE"})
			fmt.Println("ENTERED KEY NOT FOUND IN THE TABLE")
			return
		}
		// if key == "quantity" && value.(int) < 0 {
		// 	json.NewEncoder(w).Encode(map[string]string{"response": "QUANTITY CANNOT BE NEGATIVE"})
		// 	return
		// }
		query = query + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE product_id=" + args["id"]
		_, erro := helpers.RunQuery(query, w)
		helpers.HandleError("ERROR IN RUNNING UPDATE", erro, w)

		if erro != nil {
			json.NewEncoder(w).Encode(erro)
			return
		}
		query = "UPDATE inventory SET "

	}
	json.NewEncoder(w).Encode(map[string]string{"response": "UPDATE ON INVENTORY DONE"})

}

func contains(word string) bool {
	for _, elem := range key_elements {
		if word == elem {
			return true
		}
	}
	return false
}
