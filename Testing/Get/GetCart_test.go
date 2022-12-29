package Testing

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sabhari/product_catlog/helpers"
)

func GetCartViaAPI(ref string, t *testing.T) []map[string]any {
	response, err := http.Get("http://localhost:8080/cart/get?ref=" + ref)
	//fmt.Println("response here ", response)
	helpers.HandleTestError("httpGetError", err, t)

	response_json := []map[string]any{}
	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetCart(t *testing.T) {

	reference := "bd66cd83-9a1a-436a-a639-714642489c6d"
	response := GetCartViaAPI(reference, t)
	//fmt.Println("on test", response)
	_, ok := response[0]["reference"]
	//fmt.Println(ok)
	if !ok {
		t.Errorf("Expected Response: %v, Got Response: %v", "A Valid Product Map", response)
	}

	reference = "bd66cd83-9a1a-436a-a639-714642475136"
	response = GetCartViaAPI(reference, t)
	//fmt.Println("response here is", response, len(response))
	//message, ok := response[0]["response"]
	if len(response) != 0 {
		t.Errorf("Expected Response: %v, Got Response: %v", "[]", response)
	}
}
