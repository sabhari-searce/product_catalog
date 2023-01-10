package shorthandhelpers

import (
	service "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetCategoryHelper(id int) int {

	list_of_category := []typedefs.Category{}

	service.GetCategoryBL(map[string]string{"id": string(id)}, list_of_category)

	if len(list_of_category) == 0 {
		return 404
	}
	return 200

}
