package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func HandleError(errorString string, err error, w http.ResponseWriter) {
	if err != nil {
		error_string := fmt.Sprintln(errorString+" : ", err)
		fmt.Println(error_string)
		json.NewEncoder(w).Encode(error_string)

		file, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		log.SetOutput(file)
		log.Println(error_string)
	}
}
