package typedefs

type GetCartStruct struct {
	Reference     string            `json:"reference"`
	ProductID     int               `json:"product_id"`
	ProductName   string            `json:"product_name"`
	Specification map[string]string `json:"specification"`
	CategoryName  string            `json:"category_name"`
	Price         float32           `json:"price"`
	Quantity      int               `json:"quantity"`
}
