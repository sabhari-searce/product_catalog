package helpers

import (
	"database/sql"
	"net/http"
)

func RunQuery(query string, w http.ResponseWriter, v ...any) (*sql.Rows, error) {
	db := ConnectToDB()
	//fmt.Println("query is ", query)
	stmt, err := db.Prepare(query)

	HandleError("DB Prepare error", err, w)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	//var res_rows *sql.Rows
	res_rows, err := stmt.Query(v...)
	HandleError("Execution error", err, w)
	return res_rows, err
}
