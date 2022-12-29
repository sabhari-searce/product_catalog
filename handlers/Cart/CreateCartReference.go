package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sabhari/product_catlog/helpers"
)

func CreateCartReference(w http.ResponseWriter, r *http.Request) {
	reference := uuid.New()
	query := "INSERT INTO cart_reference VALUES($1,$2)"

	_, err := helpers.RunQuery(query, w, reference, time.Now())
	helpers.HandleError("Error in creating reference ", err, w)

	if err != nil {
		helpers.ResponseWriteToScreen(err, "ERROR IN CREATING REFERENCE", w)
	} else {
		response := fmt.Sprintln("Reference succesfully created and the reference is ", reference)
		helpers.ResponseWriteToScreen(err, response, w)
	}

}
