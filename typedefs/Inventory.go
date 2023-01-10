package typedefs

import "github.com/gookit/validate"

type Inventory struct {
	ProductID int `validate:"required|int|min:1"`
	Quantity  int `validate:"required|int|min:1"`
}

func (f Inventory) Messages() map[string]string {
	return validate.MS{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}
