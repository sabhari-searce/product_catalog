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

func AddProduct(w http.ResponseWriter, r *http.Request) {
	product := typedefs.Product{}
	json.NewDecoder(r.Body).Decode(&product)
	v := validate.Struct(product)

	if v.Validate() { // validate ok
		spec, err := json.Marshal(product.Specification)
		helpers.HandleError(rsp.Marshal_error, err)

		err = service.AddProductBL(product, spec)
		if err != nil {
			helpers.HandleError(rsp.ProductInErr, err)
			helpers.ResponseWriteToScreen(err, rsp.ProductInErr, w)

		} else {
			helpers.HandleError(rsp.ProductInDone, err)
			helpers.ResponseWriteToScreen(err, rsp.ProductInDone, w)
		}
	} else {
		fmt.Println(v.Errors) // all error messages
	}

}
