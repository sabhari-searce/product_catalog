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

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM category WHERE category_id = $1"
	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	helpers.HandleError("Conversion Error", err, w)

	found, _ := shorthandhelpers.GetCategoryHelper(id, w)

	if found == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "DATA NOT FOUND FOR DELETING!!"})
		return
	}

	_, err = helpers.RunQuery(query, w, args["id"])

	helpers.HandleError("Error while deleting element", err, w)

	json.NewEncoder(w).Encode(map[string]string{"response": "DELETE ON CATEGORY DONE!!"})
	fmt.Println(map[string]string{"response": "DELETE ON CATEGORY DONE!!"})

}
