package filter

import (
	"github.com/anujpradhaan/mytheresa/types"
)

// To maintain a registry of all the Filters.
// Maintaining a registry allows us to extend our design of code without affecting other filters.
var filters []*types.Filter

func registerRule(filter *types.Filter) {
	if filter != nil {
		filters = append(filters, filter)
	}
}

// ApplyToProducts applies the filters to all the products
func ApplyToProducts(products []types.Product, params types.FilteringParams) []types.Product {
	for _, filter := range filters {
		products = filter.Apply(products, params)
	}
	return products
}
