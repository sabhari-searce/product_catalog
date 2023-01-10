package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	rsp "github.com/sabhari/product_catlog/Response"
	service "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	list_of_category := []typedefs.Category{}

	service.GetCategoryBL(args, list_of_category)

	if len(list_of_category) == 0 {
		json.NewEncoder(w).Encode(rsp.GetProductErr)
		helpers.HandleError(rsp.GetProductErr["response"], nil)
		return
	}

	err := json.NewEncoder(w).Encode(list_of_category[0])
	helpers.HandleError(rsp.CategoryGetErr, err)
	fmt.Println(list_of_category[0])

}
