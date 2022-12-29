package Console_Interface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Inventory() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'Inventory' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Insert\n2.Read\n3.Update\n4.Delete\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertInventory()
	} else if choice == 2 {
		ReadInventory()
	} else if choice == 3 {
		UpdateInventory()
	} else if choice == 4 {
		DeleteInventory()
	}
}

func InsertInventory() {
	fmt.Println("Please enter the valid product id")
	var product_id int
	_, err := fmt.Scanf("%d", &product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Quantity for product")
	var quantity int
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		fmt.Println(err)
	}

	own_data := fmt.Sprintf("{\"product_id\":%v,\"quantity\":%v}", product_id, quantity)

	byte_data := []byte(own_data)

	_, err = http.Post("http://localhost:8080/inventory/add", "application/json", bytes.NewBuffer(byte_data))
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

func ReadInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8080/inventory/get/" + product_id)
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

func UpdateInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
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
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/inventory/update/%v", product_id), request_body)
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

func DeleteInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/inventory/delete/%v", product_id), nil)
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
