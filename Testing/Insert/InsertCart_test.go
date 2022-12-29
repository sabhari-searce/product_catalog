package Testing

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInsertCart(t *testing.T) {
	//data := map[string]any{"field1": "value1", "field2": "value2"}

	data := []byte(``)

	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Post("http://localhost:8080/cart/add?ref=bd66cd83-9a1a-436a-a639-714642489c6d&product=4&quantity=2", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	//fmt.Printf("%v", string(body))
	//dummy := "Insert to Product"
	//fmt.Printf("%v", []byte("\"Insert to category done\"\n"))
	if string(body) != "\"THE PRODUCT HAS BEEN ADDED TO CART!\"\n" {
		t.Errorf("Expected response body 'THE PRODUCT HAS BEEN ADDED TO CART!', got '%s'", string(body))
	}

}
