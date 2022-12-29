package shorthandhelpers

import (
	"encoding/json"
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetProductHelper(id int, w http.ResponseWriter) (int, []typedefs.Product) {
	query := "SELECT * FROM product WHERE product_id=$1"
	rows, err := helpers.RunQuery(query, w, id)
	helpers.HandleError("Error in getting product", err, w)
	rows.Scan()
	list_of_product := []typedefs.Product{}

	for rows.Next() {
		new_product := typedefs.Product{}
		spec_json := ""
		err := rows.Scan(&new_product.Product_ID, &new_product.Name, &spec_json, &new_product.SKU, &new_product.CategoryID, &new_product.Price)
		helpers.HandleError("Error in rows next", err, w)
		json.Unmarshal([]byte(spec_json), &new_product.Specification)
		list_of_product = append(list_of_product, new_product)
	}

	if len(list_of_product) == 0 {
		return 404, list_of_product
	}
	return 200, list_of_product

}
