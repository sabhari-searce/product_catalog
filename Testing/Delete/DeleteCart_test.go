package Testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func TestDeleteCart(t *testing.T) {
	//Valid Product ID
	product_id := 2
	DeleteCartViaAPI(product_id, "DELETE ON CART DONE!!", t)

	//Invalid Product Id
	product_id = 100
	DeleteCartViaAPI(product_id, "DATA NOT FOUND FOR DELETING!!", t)

}

func DeleteCartViaAPI(product_id int, expected_response string, t *testing.T) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/cart/delete?ref=bd66cd83-9a1a-436a-a639-714642489c6d&product_id=%v", product_id), nil)
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
