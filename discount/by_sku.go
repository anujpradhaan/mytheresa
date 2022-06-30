package discount

import (
	"fmt"
	"github.com/anujpradhaan/mytheresa/types"
)

const (
	SkuID3              = "000003"
	DiscountBySKURuleID = "sku-discount"
)

// Define discounts of each category.
var skuWiseDiscount = map[string]int{
	SkuID3: 15, // SKU 000003 has 15% discount
	// We could easily extend it for other SKU's
	// Example :If we want 000001 to have 10% discount, we could just add on entry here
	// and rest is taken care by the code below
}

// Register the discount by sku rule with the discount registry
func init() {
	registerDiscountRule(
		&types.DiscountRule{
			ID:            DiscountBySKURuleID,
			ApplyDiscount: GetDiscountAmountForSKU(),
		},
	)
}

func GetDiscountAmountForSKU() func(types.Product, int) (*types.Discount, error) {
	return func(product types.Product, _ int) (*types.Discount, error) {
		discountPercentage, ok := skuWiseDiscount[product.SKU]
		if !ok {
			return &types.Discount{
				OriginalAmount: product.Price,
			}, nil
		}
		return &types.Discount{
			OriginalAmount:     product.Price,
			DiscountAmount:     discountPercentage * product.Price / 100,
			DiscountPercentage: fmt.Sprintf("%d", discountPercentage),
		}, nil
	}
}
