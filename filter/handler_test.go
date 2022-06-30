package filter

import (
	"github.com/anujpradhaan/mytheresa/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilters(t *testing.T) {
	// Table Unit tests
	tests := []struct {
		name                     string
		inputProducts            []types.Product
		inputFilter              types.FilteringParams
		expectedNumberOfProducts int
		expectedProducts         []types.Product
	}{
		{
			name: "Expect single product of sandals category",
			inputProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
			},
			inputFilter: types.FilteringParams{
				Category: "sandals",
			},
			expectedNumberOfProducts: 1,
			expectedProducts: []types.Product{
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
			},
		},
		{
			name: "Expect 2 products for price less than 30000",
			inputProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
				{
					SKU:      "00005",
					Name:     "Leather belt",
					Category: "belts",
					Price:    35000,
				},
			},
			inputFilter: types.FilteringParams{
				PriceLessThan: 30000,
			},
			expectedNumberOfProducts: 2,
			expectedProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
			},
		},
		{
			name: "Expect 3 products for price less than or equal 30000",
			inputProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
				{
					SKU:      "00005",
					Name:     "Leather belt",
					Category: "belts",
					Price:    30000,
				},
			},
			inputFilter: types.FilteringParams{
				PriceLessThan: 30000,
			},
			expectedNumberOfProducts: 3,
			expectedProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
				{
					SKU:      "00005",
					Name:     "Leather belt",
					Category: "belts",
					Price:    30000,
				},
			},
		},
		{
			name: "Expect 1 product for price less than or equal 30000 and category boots",
			inputProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
				{
					SKU:      "00005",
					Name:     "Leather belt",
					Category: "belts",
					Price:    30000,
				},
			},
			inputFilter: types.FilteringParams{
				PriceLessThan: 30000,
				Category:      "boots",
			},
			expectedNumberOfProducts: 1,
			expectedProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
			},
		},
		{
			name: "Expect no product for price less than 20000",
			inputProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
				{
					SKU:      "00005",
					Name:     "Leather belt",
					Category: "belts",
					Price:    30000,
				},
			},
			inputFilter: types.FilteringParams{
				PriceLessThan: 20000,
			},
			expectedNumberOfProducts: 0,
			expectedProducts:         []types.Product{},
		},
		{
			name: "Expect no product for caps category",
			inputProducts: []types.Product{
				{
					SKU:      "00003",
					Name:     "Leather boots",
					Category: "boots",
					Price:    25000,
				},
				{
					SKU:      "00004",
					Name:     "Leather sandals",
					Category: "sandals",
					Price:    25000,
				},
				{
					SKU:      "00005",
					Name:     "Leather belt",
					Category: "belts",
					Price:    30000,
				},
			},
			inputFilter: types.FilteringParams{
				Category: "caps",
			},
			expectedNumberOfProducts: 0,
			expectedProducts:         []types.Product{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outputProducts := ApplyToProducts(test.inputProducts, test.inputFilter)
			assert.Equal(t, test.expectedNumberOfProducts, len(outputProducts))
			assert.Equal(t, test.expectedProducts, outputProducts)
		})
	}
}
