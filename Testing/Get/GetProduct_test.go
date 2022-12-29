package Testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func GetProductViaAPI(id int, t *testing.T) map[string]string {
	response, err := http.Get("http://localhost:8080/product/get/" + fmt.Sprint(id))
	helpers.HandleTestError("httpGetError", err, t)

	response_json := map[string]string{}
	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetProduct(t *testing.T) {

	// Valid product_id
	product_id := 1
	response := GetProductViaAPI(product_id, t)
	_, ok := response["product_id"]
	if !ok {
		t.Errorf("Expected Response: %v, Got Response: %v", "A Valid Product Map", response)
	}

	product_id = 500
	response = GetProductViaAPI(product_id, t)
	message, ok := response["response"]
	if !ok || message != "NO DATA FOUND!" {
		t.Errorf("Expected Response: %v, Got Response: %v", "NO DATA FOUND!", message)
	}
}
