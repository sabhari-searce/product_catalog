package Console_Interface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Product() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'product' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Insert\n2.Read with product id\n3.Update\n4.Delete\n5.Short Description with page no\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertProduct()
	} else if choice == 2 {
		ReadProduct()
	} else if choice == 3 {
		UpdateProduct()
	} else if choice == 4 {
		DeleteProduct()
	} else if choice == 5 {
		ReadShortProduct()
	}
}

func InsertProduct() {
	fmt.Println("Please enter the valid product id")
	var product_id int
	_, err := fmt.Scanf("%d", &product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the product name")
	var name string
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the specification key")
	var key string
	_, err = fmt.Scanln(&key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the specification value")
	var value string
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the SKU")
	var sku int
	_, err = fmt.Scanf("%d", &sku)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Category id")
	var category_id int
	_, err = fmt.Scanf("%d", &category_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Price")
	var price float64
	_, err = fmt.Scanf("%f", &price)
	if err != nil {
		fmt.Println(err)
	}

	own_data := fmt.Sprintf("{\"product_id\":%v,\"name\":\"%v\",\"specification\": {\"%v\":\"%v\"},\"sku\":\"%v\",\"category_id\":%v,\"price\":%v}", product_id, name, key, value, sku, category_id, price)

	byte_data := []byte(own_data)

	_, err = http.Post("http://localhost:8080/product/add", "application/json", bytes.NewBuffer(byte_data))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}

}

func ReadProduct() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8080/product/get/" + product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}

}

func ReadShortProduct() {
	fmt.Println("Please enter the page number to get the product")
	var page string
	_, err := fmt.Scanln(&page)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8080/product/getpage/" + page)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}
}

func UpdateProduct() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the key to be updated")
	var key string
	_, err = fmt.Scanln(&key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the value to be updated")
	var value string
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]any{key: value}
	byte_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	request_body := bytes.NewBuffer(byte_data)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/product/update/%v", product_id), request_body)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Update done succesfully")

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}
}

func DeleteProduct() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/product/delete/%v", product_id), nil)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete done succesfully")

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}

}
