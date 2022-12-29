package Testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func GetInventoryViaAPI(id int, t *testing.T) map[string]any {
	response, err := http.Get("http://localhost:8080/inventory/get/" + fmt.Sprint(id))
	helpers.HandleTestError("httpGetError", err, t)

	response_json := map[string]any{}

	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetInventory(t *testing.T) {

	product_id := 1
	response := GetInventoryViaAPI(product_id, t)
	//fmt.Println(response)
	_, ok := response["Product_ID"]
	//fmt.Println(ok)
	if !ok {
		t.Errorf("Expected Response: %v, Got Response: %v", "A Valid Product Map", response)
	}

	product_id = 500
	response = GetInventoryViaAPI(product_id, t)
	message, ok := response["response"]
	if !ok || message != "NO DATA FOUND!" {
		t.Errorf("Expected Response: %v, Got Response: %v", "NO DATA FOUND!", message)
	}
}
