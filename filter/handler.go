package filter

import (
	"github.com/anujpradhaan/mytheresa/types"
)

// To maintain a registry of all the Filters on the
var filters []*types.Filter

func registerRule(filter *types.Filter) {
	if filter != nil {
		filters = append(filters, filter)
	}
}

func ApplyToProducts(products []types.Product, params types.FilteringParams) []types.Product {
	for _, filter := range filters {
		products = filter.Apply(products, params)
	}
	return products
}
