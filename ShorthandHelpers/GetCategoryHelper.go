package shorthandhelpers

import (
	"net/http"

	"github.com/sabhari/product_catlog/helpers"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCategoryHelper(id int, w http.ResponseWriter) (int, []typedefs.Category) {
	query := "SELECT * FROM category WHERE category_id=$1"
	rows, err := helpers.RunQuery(query, w, id)
	helpers.HandleError("Error in getting category", err, w)
	rows.Scan()
	list_of_category := []typedefs.Category{}

	for rows.Next() {
		new_category := typedefs.Category{}
		err := rows.Scan(&new_category.CategoryID, &new_category.Name)
		helpers.HandleError("Error in rows next", err, w)
		list_of_category = append(list_of_category, new_category)
	}

	if len(list_of_category) == 0 {
		return 404, list_of_category
	}
	return 200, list_of_category

}
