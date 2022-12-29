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

var key_elements []string = []string{"name"}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category map[string]any
	json.NewDecoder(r.Body).Decode(&category)
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	helpers.HandleError("Eror when converting string to int", err, w)

	status, _ := shorthandhelpers.GetCategoryHelper(id, w)
	if status == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED ID NOT FOUND FOR UPDATING"})
		return
	}

	query := "UPDATE category SET "

	for key, value := range category {

		if !contains(key) {
			json.NewEncoder(w).Encode(map[string]string{"response": "ENTERED KEY NOT FOUND IN THE TABLE"})
			fmt.Println("ENTERED KEY NOT FOUND IN THE TABLE")
			return
		}
		query = query + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE category_id=" + args["id"]
		_, erro := helpers.RunQuery(query, w)
		helpers.HandleError("ERROR IN RUNNING UPDATE", erro, w)
		// res_string := fmt.Sprintln(key, " UPDATED ON CATEGORY")
		// helpers.ResponseWriteToScreen(erro, res_string, w)
		if erro != nil {
			json.NewEncoder(w).Encode(erro)
			return
		}
		query = "UPDATE category SET "

	}
	json.NewEncoder(w).Encode(map[string]string{"response": "UPDATED ON CATEGORY DONE!"})

}

func contains(word string) bool {
	for _, elem := range key_elements {
		if word == elem {
			return true
		}
	}
	return false
}
