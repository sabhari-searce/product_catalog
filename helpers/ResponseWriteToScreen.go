package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseWriteToScreen(err error, query_operation string, w http.ResponseWriter) {
	if err == nil {
		Response_message := fmt.Sprintf("%v DONE SUCCESFULLY", query_operation)
		fmt.Println(Response_message)
		json.NewEncoder(w).Encode(Response_message)
	}

}
