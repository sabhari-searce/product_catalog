package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM category WHERE category_id = $1"
	args := mux.Vars(r)
	list_of_category := []typedefs.Category{}

	rows, err := helpers.RunQuery(query, w, args["id"])
	rows.Scan()
	helpers.HandleError("Error in getting Category", err, w)

	for rows.Next() {
		new_category := typedefs.Category{}
		err := rows.Scan(&new_category.CategoryID, &new_category.Name)
		helpers.HandleError("Error in rows next", err, w)
		list_of_category = append(list_of_category, new_category)
	}

	if len(list_of_category) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"response": "NO DATA FOUND!"})
		return
	}

	err = json.NewEncoder(w).Encode(list_of_category[0])
	helpers.HandleError("Error in getting category", err, w)
	fmt.Println(list_of_category[0])

}
