package typedefs

import "github.com/gookit/validate"

type Cart struct {
	Reference string `validate:"required|minLen:5|string"`
	ProductID int    `validate:"required|int|min:1"`
	Quantity  int    `validate:"required|int|min:1"`
}

func (f Cart) Messages() map[string]string {
	return validate.MS{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}
