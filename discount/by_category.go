package discount

import (
	"fmt"
	"github.com/anujpradhaan/mytheresa/types"
)

const (
	CategoryBoots  = "boots"
	CategoryRuleID = "category-discount"
)

// Define discounts of each category.
var categoryWiseDiscount = map[string]int{
	CategoryBoots: 30, // boots category has 30% discount
}

// Register the discount category rule with the discount registry
func init() {
	registerDiscountRule(
		&types.DiscountRule{
			ID:            CategoryRuleID,
			ApplyDiscount: GetDiscountAmountForCategory(),
		},
	)
}

func GetDiscountAmountForCategory() func(types.Product, int) (*types.Discount, error) {
	return func(product types.Product, _ int) (*types.Discount, error) {
		discountPercentage, ok := categoryWiseDiscount[product.Category]
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
