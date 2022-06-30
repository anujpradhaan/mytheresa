package types

// DiscountRule is a general Discount type, If we
// wish to create a new category of discount rules. We should use this type
type DiscountRule struct {
	ID            string
	ApplyDiscount func(product Product, originalPrice int) (*Discount, error)
}

// Discount represents a discount applied on a given amount together with the applicable percentage
type Discount struct {
	OriginalAmount     int
	DiscountAmount     int
	DiscountPercentage string
}
