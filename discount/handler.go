package discount

import (
	"github.com/anujpradhaan/mytheresa/types"
)

// To maintain a registry of all the Discount Rules
var discountRules []*types.DiscountRule

func registerDiscountRule(discount *types.DiscountRule) {
	if discount != nil {
		discountRules = append(discountRules, discount)
	}
}

// Calculate runs all the discount rules from the registry and return the biggest one
func Calculate(product types.Product) (*types.Discount, error) {
	var biggestDiscount *types.Discount

	for _, discountRule := range discountRules {

		intermediateDiscount, err := discountRule.ApplyDiscount(product, product.Price)
		if err != nil {
			//TODO add a logging library
			return nil, err
		}

		if biggestDiscount == nil || biggestDiscount.DiscountAmount < intermediateDiscount.DiscountAmount {
			biggestDiscount = intermediateDiscount
		}
	}

	// No discount can be applied for the product
	if biggestDiscount == nil {
		biggestDiscount = &types.Discount{
			OriginalAmount: product.Price,
		}
	}
	return biggestDiscount, nil
}
