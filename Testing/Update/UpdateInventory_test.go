package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func TestUpdateInventory(t *testing.T) {
	//valid update operation
	product_id := 2
	update_json_req_body_map := map[string]any{
		"quantity": 15,
	}

	CheckUpdateEndpointInventory(product_id, update_json_req_body_map, "UPDATE ON INVENTORY DONE", t)

	//invalid product_id
	delete(update_json_req_body_map, "product_id")
	product_id = 100
	CheckUpdateEndpointInventory(product_id, update_json_req_body_map, "ENTERED ID NOT FOUND FOR UPDATING", t)

}

func CheckUpdateEndpointInventory(product_id int, update_json_req_body_map map[string]any, expected_response string, t *testing.T) {
	json_product, err := json.Marshal(update_json_req_body_map)
	helpers.HandleTestError("jsonMarshalError", err, t)

	request_body := bytes.NewBuffer(json_product)
	//fmt.Printf("http://localhost:8080/product/update/%v", product_id)
	//fmt.Println(string(json_product))
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/inventory/update/%v", product_id), request_body)
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
