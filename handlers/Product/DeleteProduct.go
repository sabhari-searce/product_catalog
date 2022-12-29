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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM product WHERE product_id = $1"
	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	helpers.HandleError("Conversion Error", err, w)

	found, _ := shorthandhelpers.GetProductHelper(id, w)

	if found == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "DATA NOT FOUND FOR DELETING!!"})
		return
	}

	_, err = helpers.RunQuery(query, w, args["id"])

	helpers.HandleError("Error while deleting element", err, w)

	//helpers.ResponseWriteToScreen(err, "DELETE ON PRODUCT", w)
	json.NewEncoder(w).Encode(map[string]string{"response": "DELETE ON PRODUCT DONE"})
	fmt.Println(map[string]string{"response": "DELETE ON PRODUCT DONE"})

}
