package types

type FilteringParams struct {
	Category      string `form:"category"`
	PriceLessThan int    `form:"priceLessThan"`
}

type Filter struct {
	Name  string
	Apply func(data []Product, params FilteringParams) []Product
}
