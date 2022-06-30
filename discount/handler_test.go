package discount

import (
	"github.com/anujpradhaan/mytheresa/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiscountEngine(t *testing.T) {
	// Table Unit tests
	tests := []struct {
		name             string
		inputProduct     types.Product
		expectedDiscount *types.Discount
	}{
		{
			name: "Expect a 30% discount as boots has higher discount",
			inputProduct: types.Product{
				SKU:      SkuID3,
				Name:     "Leather boots",
				Category: "boots",
				Price:    25000,
			},
			expectedDiscount: &types.Discount{
				OriginalAmount:     25000,
				DiscountPercentage: "30",
				DiscountAmount:     7500,
			},
		},
		{
			// This might not be a real work scenario as SKU's might
			// be different for every product
			// But it's important to see the behaviour
			name: "Expect a 15% discount as towels has 15% discount for SKUID3",
			inputProduct: types.Product{
				SKU:      SkuID3,
				Name:     "Cotton towels",
				Category: "towels",
				Price:    25000,
			},
			expectedDiscount: &types.Discount{
				OriginalAmount:     25000,
				DiscountPercentage: "15",
				DiscountAmount:     3750,
			},
		},
		{
			name: "Expect 30 discount for non 000003 SKU and boots category",
			inputProduct: types.Product{
				SKU:      "12345",
				Name:     "Leather boots Armani",
				Category: "boots",
				Price:    25000,
			},
			expectedDiscount: &types.Discount{
				OriginalAmount:     25000,
				DiscountAmount:     7500,
				DiscountPercentage: "30",
			},
		},
		{
			name: "Expect no discount",
			inputProduct: types.Product{
				SKU:      "12346",
				Name:     "Leather Belt",
				Category: "belt",
				Price:    25000,
			},
			expectedDiscount: &types.Discount{
				OriginalAmount: 25000,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			discount, err := Calculate(test.inputProduct)
			require.NoError(t, err)
			assert.Equal(t, test.expectedDiscount, discount)
		})
	}
}
