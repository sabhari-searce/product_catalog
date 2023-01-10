package typedefs

import "github.com/gookit/validate"

type Category struct {
	CategoryID int    `validate:"required|int|min:1"`
	Name       string `validate:"required|minLen:3|string"`
}

func (f Category) Messages() map[string]string {
	return validate.MS{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}
