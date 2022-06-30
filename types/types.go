package types

// Product is to represent a product we would want to sell.
type Product struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"` //TODO : Check if int is the suitable type here
}

// ProductData
type ProductData struct {
	Products []Product `json:"products"`
}

// ProductResponse is to represent the response returned to client
type ProductResponse struct {
	SKU            string              `json:"sku"`
	Name           string              `json:"name"`
	Category       string              `json:"category"`
	ProductPricing ProductPricingModel `json:"price"` //TODO : Check if int is the suitable type here
}

// ProductPricingModel to represent the pricing model of a product
type ProductPricingModel struct {
	OriginalPrice      int     `json:"original"` //TODO check if we can set validations
	FinalPrice         int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"` // TODO, how to set default value
}
