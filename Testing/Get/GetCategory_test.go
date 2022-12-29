package Testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func GetCategoryViaAPI(id int, t *testing.T) map[string]string {
	response, err := http.Get("http://localhost:8080/category/get/" + fmt.Sprint(id))
	helpers.HandleTestError("httpGetError", err, t)

	response_json := map[string]string{}
	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetCategory(t *testing.T) {

	category_id := 1
	response := GetCategoryViaAPI(category_id, t)
	//fmt.Println(response)
	_, ok := response["category_id"]
	//fmt.Println(ok)
	if !ok {
		t.Errorf("Expected Response: %v, Got Response: %v", "A Valid Product Map", response)
	}

	category_id = 500
	response = GetCategoryViaAPI(category_id, t)
	message, ok := response["response"]
	if !ok || message != "NO DATA FOUND!" {
		t.Errorf("Expected Response: %v, Got Response: %v", "NO DATA FOUND!", message)
	}
}
