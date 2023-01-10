package shorthandhelpers

import (
	services "github.com/sabhari/product_catlog/Services"
	"github.com/sabhari/product_catlog/typedefs"
)

func GetProductHelper(id int) int {

	list_of_product := []typedefs.ProductDesc{}

	services.GetProductBL(list_of_product, string(id))

	if len(list_of_product) == 0 {
		return 404
	}
	return 200

}
