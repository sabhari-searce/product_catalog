package typedefs

type ProductDesc struct {
	Product_ID    int               `json:"product_id"`
	Name          string            `json:"name"`
	Specification map[string]string `json:"specification"`
	SKU           string            `json:"sku"`
	Category_name string            `json:"category_name"`
	Price         float32           `json:"price"`
}
