package services

import (
	"fmt"

	query "github.com/sabhari/product_catlog/Queries"
	rsp "github.com/sabhari/product_catlog/Response"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func AddCategoryBL(category typedefs.Category) int {
	_, err := helpers.RunQuery(query.AddCategory, category.CategoryID, category.Name)
	if err != nil {
		return 404
	}
	return 200
}

func GetCategoryBL(args map[string]string, list_of_category []typedefs.Category) {
	rows, err := helpers.RunQuery(query.GetCategory, args["id"])
	rows.Scan()
	helpers.HandleError(rsp.CategoryGetErr, err)

	for rows.Next() {
		new_category := typedefs.Category{}
		err := rows.Scan(&new_category.CategoryID, &new_category.Name)
		helpers.HandleError(rsp.GetRowErr, err)
		list_of_category = append(list_of_category, new_category)
	}
}

func DeleteCategoryBL(args map[string]string) int {
	_, err := helpers.RunQuery(query.DeleteCategory, args["id"])
	if err != nil {
		return 404
	}
	return 200

}

var key_elements_cat []string = []string{"name"}

func containsCategory(word string) bool {
	for _, elem := range key_elements_cat {
		if word == elem {
			return true
		}
	}
	return false
}

func UpdateCategoryBL(category map[string]any, args map[string]string) int {
	queryexe := query.UpdateCategory
	for key, value := range category {

		if !contains(key) {
			return 403
		}
		queryexe = queryexe + key + "='" + fmt.Sprintf("%v", value) + "'" + " WHERE category_id=" + args["id"]
		_, erro := helpers.RunQuery(queryexe)
		helpers.HandleError("ERROR IN RUNNING UPDATE", erro)
		// res_string := fmt.Sprintln(key, " UPDATED ON CATEGORY")
		// helpers.ResponseWriteToScreen(erro, res_string, w)
		if erro != nil {
			return 404
		}
		queryexe = query.UpdateCategory
	}
	return 200
}
