package typedefs

import "github.com/gookit/validate"

type Product struct {
	Product_ID    int               `validate:"required|int|min:1"`
	Name          string            `validate:"required|minLen:5|string"`
	Specification map[string]string `validate:required`
	SKU           string            `validate:"required|string"`
	CategoryID    int               `validate:"required|int|min:1"`
	Price         float32           `validate:"required|int|min:1"`
}

func (f Product) Messages() map[string]string {
	return validate.MS{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}
