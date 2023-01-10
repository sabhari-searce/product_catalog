package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gookit/validate"
	rsp "github.com/sabhari/product_catlog/Response"
	service "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	category := typedefs.Category{}
	json.NewDecoder(r.Body).Decode(&category)

	v := validate.Struct(category)
	// v := validate.New(u)
	if v.Validate() { // validate ok
		category_id := category.CategoryID
		if category_id <= 0 {
			json.NewEncoder(w).Encode(rsp.CategoryIdErr)
			helpers.HandleError(rsp.CategoryIdErr, nil)
		}

		category_name := category.Name
		if category_name == "" {
			json.NewEncoder(w).Encode(rsp.CategoryNameErr)
			helpers.HandleError(rsp.CategoryNameErr, nil)
		}

		found := service.AddCategoryBL(category)
		if found == 404 {
			json.NewEncoder(w).Encode(rsp.CategoryInErr)
			helpers.HandleError(rsp.CategoryInErr, nil)
		} else if found == 200 {
			json.NewEncoder(w).Encode(rsp.CategoryIn)
			helpers.HandleError(rsp.CategoryIn, nil)
		}

	} else {
		fmt.Println(v.Errors) // all error messages
	}

}
