package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	query "github.com/sabhari/product_catlog/Queries"
	rsp "github.com/sabhari/product_catlog/Response"
	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddProductBL(product typedefs.Product, spec []byte) error {
	product_id := product.Product_ID
	if product_id <= 0 {
		helpers.HandleError(rsp.ProductIdErr, nil)

	}

	product_name := product.Name
	if product_name == "" {
		helpers.HandleError(rsp.ProductNameErr, nil)
	}

	category := product.CategoryID
	if category <= 0 {
		helpers.HandleError(rsp.CategoryIdErr, nil)
	}

	price := product.Price
	if price <= 0.0 {
		helpers.HandleError(rsp.PriceErr, nil)
	}
	_, err := helpers.RunQuery(query.AddProduct, product_id, product_name, spec, product.SKU, category, product.Price)
	if err != nil {
		helpers.HandleError("Error in inserting", err)
		return err
	}
	return nil
}

func GetProductBL(list_of_product []typedefs.ProductDesc, id string) {
	check_id, err := strconv.Atoi(id)
	if err != nil || check_id <= 0 {
		helpers.HandleError(rsp.ProductIdErr, nil)
	}
	rows, err := helpers.RunQuery(query.GetProduct, id)
	rows.Scan()
	helpers.HandleError(rsp.ProductGetErr, err)

	for rows.Next() {
		new_product := typedefs.ProductDesc{}
		spec_json := ""
		err := rows.Scan(&new_product.Product_ID, &new_product.Name, &spec_json, &new_product.SKU, &new_product.Category_name, &new_product.Price)
		helpers.HandleError(rsp.GetRowErr, err)
		json.Unmarshal([]byte(spec_json), &new_product.Specification)
		list_of_product = append(list_of_product, new_product)
	}
}

func DeleteProductBL(id int) {
	_, err := helpers.RunQuery(query.DeleteProduct, id)

	helpers.HandleError("Error while deleting element", err)
}

func GetShortProductBL(args map[string]string) ([]typedefs.ShortProduct, int) {
	list_of_product := []typedefs.ShortProduct{}

	rows, err := helpers.RunQuery(query.GetProductShort)
	rows.Scan()
	helpers.HandleError(rsp.ProductGetErr, err)

	for rows.Next() {
		new_product := typedefs.ShortProduct{}
		err := rows.Scan(&new_product.Product_ID, &new_product.Name, &new_product.Price)
		helpers.HandleError(rsp.GetRowErr, err)
		list_of_product = append(list_of_product, new_product)
	}

	code := 200
	max_length_obj := len(list_of_product)
	argument, err := strconv.Atoi(args["id"])
	helpers.HandleError(rsp.AtoiErr, err)
	if argument <= 0 {
		helpers.HandleError(rsp.GetProductErr["response"], nil)
		code = 404
	}
	end_index := (argument * 20)
	start_index := end_index - 20

	//var res []typedefs.ShortProduct
	if start_index > max_length_obj {
		//panic("Requested Page not found")
		helpers.HandleError(rsp.GetProductErr["response"], nil)
		code = 404
	}
	if end_index > max_length_obj {
		//fmt.Println("Entered here")
		end_index = max_length_obj
		//fmt.Println(end_index)
	}
	res := list_of_product[start_index:end_index]
	return res, code

}

var key_elements []string = []string{"name", "specification", "sku", "category_id", "price"}

func contains(word string) bool {
	for _, elem := range key_elements {
		if word == elem {
			return true
		}
	}
	return false
}

func UpdateProductBL(product map[string]any, id string) int {
	queryexe := query.UpdateProduct

	for key, value := range product {

		if key == "specification" {
			new_value, err := json.Marshal(value.(map[string]any))
			queryexe = queryexe + key + "='" + string(new_value) + "'" + " WHERE product_id=" + id
			helpers.HandleError(rsp.Marshal_error, err)
			_, erro := helpers.RunQuery(queryexe)
			if erro != nil {
				return 404
			}
			helpers.HandleError(rsp.ProductUpErr, erro)

		} else {
			if !contains(key) {
				helpers.HandleError(rsp.ProductKeyErr, nil)
				return 403
				break
			}
			queryexe = queryexe + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE product_id=" + id
			_, erro := helpers.RunQuery(queryexe)
			helpers.HandleError(rsp.ProductUpErr, erro)

			if erro != nil {
				return 404
			}
			queryexe = query.UpdateProduct

		}

	}
	return 200
}
