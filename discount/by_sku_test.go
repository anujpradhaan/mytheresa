package discount

import (
	"github.com/anujpradhaan/mytheresa/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDiscountAmountForSKU(t *testing.T) {
	// Table Unit tests
	tests := []struct {
		name             string
		description      string
		inputProduct     types.Product
		expectedDiscount *types.Discount
	}{
		{
			name: "Expect a 15% discount for SKU 000003",
			inputProduct: types.Product{
				SKU:      SkuID3,
				Name:     "Leather boots",
				Category: "boots",
				Price:    25000,
			},
			expectedDiscount: &types.Discount{
				OriginalAmount:     25000,
				DiscountPercentage: "15",
				DiscountAmount:     3750,
			},
		},
		{
			name: "Expect no discount for non 000003 SKU",
			inputProduct: types.Product{
				SKU:      "12345",
				Name:     "Leather boots",
				Category: "sneaker",
				Price:    25000,
			},
			expectedDiscount: &types.Discount{
				OriginalAmount: 25000,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			discountFunc := GetDiscountAmountForSKU()
			discount, err := discountFunc(test.inputProduct, 0)
			require.NoError(t, err)
			assert.Equal(t, test.expectedDiscount, discount)
		})
	}
}
