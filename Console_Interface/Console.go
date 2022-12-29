package Console_Interface

import (
	"fmt"
)

func Console() {
	fmt.Println("WELCOME TO CONSOLE INTERFACE")
	fmt.Println("Select the table to perform the task")
	fmt.Printf("1.Product\n2.Category\n3.Inventory\n4.Cart\n")
	fmt.Println("Enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}

	if choice == 1 {
		Product()
	} else if choice == 2 {
		Category()
	} else if choice == 3 {
		Inventory()
	} else if choice == 4 {
		Cart()
	}

}
