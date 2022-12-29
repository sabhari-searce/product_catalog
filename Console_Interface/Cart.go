package Console_Interface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Cart() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'Cart' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Insert\n2.Read\n3.Update\n4.Delete\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertCart()
	} else if choice == 2 {
		ReadCart()
	} else if choice == 3 {
		UpdateCart()
	} else if choice == 4 {
		DeleteCart()
	}
}

func InsertCart() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the valid product id")
	var product_id string
	_, err = fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the quantity of product")
	var quantity string
	_, err = fmt.Scanln(&quantity)
	if err != nil {
		fmt.Println(err)
	}

	url := "http://localhost:8080/cart/add?ref=" + ref + "&product=" + product_id + "&quantity=" + quantity

	_, err = http.Post(url, "application/json", nil)
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

func ReadCart() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8080/cart/get?ref=" + ref)
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

func UpdateCart() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the valid product id")
	var product_id string
	_, err = fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	key := "quantity"

	fmt.Println("Please enter the quantity to be updated")
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

	url := fmt.Sprintf("http://localhost:8080/cart/update?ref=%v&product_id=%v", ref, product_id)
	//fmt.Println(url)
	req, err := http.NewRequest("POST", url, request_body)
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

func DeleteCart() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the valid product id")
	var product_id string
	_, err = fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("http://localhost:8080/cart/delete?ref=%v&product_id=%v", ref, product_id)

	req, err := http.NewRequest("GET", url, nil)
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
