package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	rsp "github.com/sabhari/product_catlog/Response"
	service "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/helpers"
)

func CreateCartReference(w http.ResponseWriter, r *http.Request) {

	reference := service.CreateCartReferenceBL()
	if reference == uuid.Nil {
		json.NewEncoder(w).Encode(rsp.CreateCartReferenceErr)
		helpers.HandleError(rsp.CreateCartReferenceErr, nil)
	} else {
		res := fmt.Sprint("Reference has been created and the reference is ", reference)
		json.NewEncoder(w).Encode(res)
		helpers.HandleError(res, nil)

	}

}
