package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	rsp "github.com/sabhari/product_catlog/Response"
	service "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/helpers"
)

func GetShortProduct(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	res, code := service.GetShortProductBL(args)
	if code == 404 {
		json.NewEncoder(w).Encode(rsp.GetProductErr)
		helpers.HandleError(rsp.GetProductErr["response"], nil)
	} else {
		json.NewEncoder(w).Encode(res)
		fmt.Println(res)
	}
}
