package filter

import (
	"github.com/anujpradhaan/mytheresa/types"
)

const (
	CategoryFilter = "category-filter"
)

func init() {
	registerRule(&types.Filter{
		Name: CategoryFilter,
		Apply: func(products []types.Product, params types.FilteringParams) []types.Product {
			//If no category provided, return all products
			if params.Category == "" {
				return products
			}

			var subProducts []types.Product
			for _, pr := range products {
				if pr.Category == params.Category {
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
