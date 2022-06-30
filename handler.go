package main

import (
	"encoding/json"
	"fmt"
	"github.com/anujpradhaan/mytheresa/discount"
	"github.com/anujpradhaan/mytheresa/filter"
	"github.com/anujpradhaan/mytheresa/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

const DefaultCurrency = "EUR"

// GetProducts serves as an entry point to API and is responsible
// for reading all the products from the JSON data file.
// It also, trims the results to max 5 as expected in the problem statement.
func GetProducts(c *gin.Context) {
	filters := types.FilteringParams{}
	if err := c.ShouldBind(&filters); err != nil {
		log.Printf("error getting query params")
		c.Error(err)
		return
	}

	// List all the products
	products, err := ListAllProducts(&filters)
	if err != nil {
		c.Error(err)
		return
	}

	productsResponse := make([]types.ProductResponse, 0)
	for _, singleProduct := range products {

		// Calculate discount for the product
		calculatedDiscount, err := discount.Calculate(singleProduct)
		if err != nil {
			log.Printf("error processing discount %v", map[string]string{
				"sku": singleProduct.SKU,
			})
			c.Error(err)
			return
		}
		// Create a response with help of calculated discount
		productsResponse = append(productsResponse, getProductResponse(singleProduct, calculatedDiscount))
	}
	// return only 5 elements list of products
	if len(productsResponse) > 5 {
		// Not applying any complex logic as the only ask in the problem
		// statement was to return max 5 elements to the response.
		productsResponse = productsResponse[:5]
	}
	c.JSON(http.StatusOK, productsResponse)
}

func getProductResponse(product types.Product, discount *types.Discount) types.ProductResponse {
	return types.ProductResponse{
		SKU:      product.SKU,
		Name:     product.Name,
		Category: product.Category,
		ProductPricing: types.ProductPricingModel{
			OriginalPrice:      discount.OriginalAmount,
			DiscountPercentage: getFormattedDiscountPercent(discount),
			FinalPrice:         discount.OriginalAmount - discount.DiscountAmount,
			Currency:           DefaultCurrency,
		},
	}
}

func getFormattedDiscountPercent(discount *types.Discount) *string {
	if discount.DiscountPercentage == "" {
		return nil
	}
	s := fmt.Sprintf("%s%%", discount.DiscountPercentage)
	return &s
}

func ListAllProducts(filters *types.FilteringParams) ([]types.Product, error) {
	var productData types.ProductData

	// Reading from file and keeping in memory as implementing a DB related approach would
	// require a DB installation beforehand which might not be possible on any other machine
	file, err := ioutil.ReadFile("products.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &productData)
	if err != nil {
		return nil, err
	}
	if filters == nil {
		return productData.Products, nil
	}

	filteredProducts := filter.ApplyToProducts(productData.Products, *filters)
	return filteredProducts, nil
}
