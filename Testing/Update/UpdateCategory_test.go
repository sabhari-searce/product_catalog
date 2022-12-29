package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func TestUpdateCategory(t *testing.T) {
	//valid update operation
	category_id := 15
	update_json_req_body_map := map[string]any{
		"name": "SMART GADGETS",
	}

	CheckUpdateEndpointCategory(category_id, update_json_req_body_map, "UPDATED ON CATEGORY DONE!", t)

	//invalid category_id
	delete(update_json_req_body_map, "category_id")
	category_id = 100
	CheckUpdateEndpointCategory(category_id, update_json_req_body_map, "ENTERED ID NOT FOUND FOR UPDATING", t)

}

func CheckUpdateEndpointCategory(category_id int, update_json_req_body_map map[string]any, expected_response string, t *testing.T) {
	json_product, err := json.Marshal(update_json_req_body_map)
	helpers.HandleTestError("jsonMarshalError", err, t)

	request_body := bytes.NewBuffer(json_product)
	//fmt.Printf("http://localhost:8080/product/update/%v", category_id)
	//fmt.Println(string(json_product))
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/category/update/%v", category_id), request_body)
	helpers.HandleTestError("httpNewRequestError", err, t)

	res, err := http.DefaultClient.Do(req)
	helpers.HandleTestError("httpDefaultClientDoError", err, t)

	var v map[string]string
	json.NewDecoder(res.Body).Decode(&v)

	//fmt.Println(v)

	if expected_response != v["response"] {
		t.Errorf("Expected: %v, Got: %v", expected_response, v["response"])
	}

}
