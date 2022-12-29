package typedefs

type Cart struct {
	Reference string `json:"reference"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
