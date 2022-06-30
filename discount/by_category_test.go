package discount

import (
	"github.com/anujpradhaan/mytheresa/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDiscountAmountForCategory(t *testing.T) {
	// Table Unit tests
	tests := []struct {
		name             string
		inputProduct     types.Product
		expectedDiscount *types.Discount
	}{
		{
			name: "Expect a 30% discount for boots category",
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
			name: "Expect no discount for sneaker category",
			inputProduct: types.Product{
				SKU:      SkuID3,
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
			discountFunc := GetDiscountAmountForCategory()
			discount, err := discountFunc(test.inputProduct, 0)
			require.NoError(t, err)
			assert.Equal(t, test.expectedDiscount, discount)
		})
	}
}
