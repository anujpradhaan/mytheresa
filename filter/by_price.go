package filter

import (
	"github.com/anujpradhaan/mytheresa/types"
)

const (
	PriceFilter = "price-filter"
)

func init() {
	registerRule(&types.Filter{
		Name: PriceFilter,
		Apply: func(products []types.Product, params types.FilteringParams) []types.Product {

			//If no price provided or negative price provided, return all products
			if params.PriceLessThan <= 0 {
				return products
			}

			var subProducts []types.Product
			for _, pr := range products {
				if pr.Price <= params.PriceLessThan {
					subProducts = append(subProducts, pr)
				}
			}

			// No product matching filtering criteria
			if len(subProducts) == 0 {
				return []types.Product{}
			}
			return subProducts
		},
	})
}
