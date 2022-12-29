package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetShortProduct(w http.ResponseWriter, r *http.Request) {
	query := "SELECT product_id,name,price FROM product"
	args := mux.Vars(r)
	list_of_product := []typedefs.ShortProduct{}

	rows, err := helpers.RunQuery(query, w)
	rows.Scan()
	helpers.HandleError("Error in getting product", err, w)

	for rows.Next() {
		new_product := typedefs.ShortProduct{}
		//spec_json := ""
		err := rows.Scan(&new_product.Product_ID, &new_product.Name, &new_product.Price)
		helpers.HandleError("Error in rows next", err, w)
		//json.Unmarshal([]byte(spec_json), &new_product.Specification)
		list_of_product = append(list_of_product, new_product)
	}

	//fmt.Println(len(list_of_product))
	max_length_obj := len(list_of_product)
	argument, err := strconv.Atoi(args["id"])
	helpers.HandleError("Cannot convert string", err, w)
	if argument == 0 {
		json.NewEncoder(w).Encode("Requested Page not found")
		return
	}
	end_index := (argument * 20)
	start_index := end_index - 20

	//var res []typedefs.ShortProduct
	if start_index > max_length_obj {
		//panic("Requested Page not found")
		json.NewEncoder(w).Encode("Requested Page not found")
		return
	}
	if end_index > max_length_obj {
		//fmt.Println("Entered here")
		end_index = max_length_obj
		//fmt.Println(end_index)
	}
	fmt.Println(start_index, end_index)
	res := list_of_product[start_index:end_index]
	err = json.NewEncoder(w).Encode(res)
	fmt.Println(res)

	//helpers.ResponseWriteToScreen(err, "FETCHING DATA FROM PRODUCT", w)

}
