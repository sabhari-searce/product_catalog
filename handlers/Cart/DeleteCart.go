package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	shorthandhelpers "github.com/sabhari/product_catlog/ShorthandHelpers"
	"github.com/sabhari/product_catlog/helpers"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM cart_item WHERE ref = $1 AND product_id = $2"
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")
	product := urlQuery.Get("product_id")

	product_id, err := strconv.Atoi(product)
	helpers.HandleError("Conversion Error", err, w)

	found, _ := shorthandhelpers.GetCartHelper(reference, product_id, w)

	if found == 404 {
		json.NewEncoder(w).Encode(map[string]string{"response": "DATA NOT FOUND FOR DELETING!!"})
		return
	}

	_, err = helpers.RunQuery(query, w, reference, product_id)

	helpers.HandleError("Error while deleting element", err, w)

	//helpers.ResponseWriteToScreen(err, "DELETE ON CART", w)
	json.NewEncoder(w).Encode(map[string]string{"response": "DELETE ON CART DONE!!"})
	fmt.Println(map[string]string{"response": "DELETE ON CART DONE!!"})

}
