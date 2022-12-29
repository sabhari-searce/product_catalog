package Testing

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInsertProduct(t *testing.T) {
	//data := map[string]any{"field1": "value1", "field2": "value2"}

	data := []byte(`{
        "product_id": 70,
        "name": "Hoodies",
        "specification": {
          "size": "40",
          "color": "Black"
        },
        "sku": "64135612345",
        "category_id": 1,
        "price": 40.29
      }`)

	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Post("http://localhost:8080/product/add", "application/json", bytes.NewBuffer(data))
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
	//fmt.Printf("%v", body)
	//dummy := "Insert to Product"
	//fmt.Printf("%v", []byte("\"Insert to Product done\"\n"))
	if string(body) != "\"Insert to Product done\"\n" {
		t.Errorf("Expected response body 'Insert to Product done', got '%s'", string(body))
	}

}
