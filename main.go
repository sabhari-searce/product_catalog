package main

import (
	"fmt"

	console_interface "github.com/sabhari/product_catlog/Console_Interface"
	"github.com/sabhari/product_catlog/app"
)

func main() {

	f := func() {
		fmt.Println("If you want to start console interface? (yes or no)")
		var reply string
		_, err := fmt.Scanln(&reply)
		if err != nil {
			fmt.Println("error in reading input!!")
		}
		if reply == "yes" {
			console_interface.Console()
		} else if reply == "no" {
			fmt.Println("Console Interface cancelled")
		}
	}

	go f()

	app.StartServer()

}
