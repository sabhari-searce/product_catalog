package Testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func TestDeleteCategory(t *testing.T) {
	//Valid Product ID
	product_id := 11
	DeleteCategoryViaAPI(product_id, "DELETE ON CATEGORY DONE!!", t)

	//Invalid Product Id
	product_id = 100
	DeleteCategoryViaAPI(product_id, "DATA NOT FOUND FOR DELETING!!", t)

}

func DeleteCategoryViaAPI(product_id int, expected_response string, t *testing.T) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/category/delete/%v", product_id), nil)
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
