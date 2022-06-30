package main

import (
	"encoding/json"
	"fmt"
	"github.com/anujpradhaan/mytheresa/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestProductsApi(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name                          string
		filterByCategory              string
		filterByAmountLessThanOrEqual int
		expectedNumberOfProducts      int
		expectedProductsResponse      []types.ProductResponse
	}{
		{
			name:                     "get all products without any filter",
			expectedProductsResponse: getExpectedProductResponseForAllProducts(),
			expectedNumberOfProducts: 5, // Expected 5 as we have to return a max of 5 items
		},
		{
			name:                     "get all products when filtered by category boots",
			filterByCategory:         "boots",
			expectedProductsResponse: getExpectedProductResponsesFilteredByCategory("boots"),
			expectedNumberOfProducts: 3, // As there are only 3 items in with category boots
		},
		{
			name:                     "get all products when filtered by category sandals",
			filterByCategory:         "sandals",
			expectedProductsResponse: getExpectedProductResponsesFilteredByCategory("sandals"),
			expectedNumberOfProducts: 1, // As there are only 1 item in with category sandals
		},
		{
			name:                          "get all products filtered by priceLessThan",
			filterByAmountLessThanOrEqual: 60000,
			expectedProductsResponse:      getExpectedProductResponsesFilteredByPrice(60000),
			expectedNumberOfProducts:      1, // As there are only 1 items in with price less than 60000
		},
		{
			name:                          "get all products filtered by category and priceLessThan",
			filterByCategory:              "caps",
			filterByAmountLessThanOrEqual: 70000,
			expectedProductsResponse:      getExpectedProductResponseFilteredByCategoryAndPrice("caps", 70000),
			expectedNumberOfProducts:      1, // As there are only 1 items in with price less than 70000 and category as caps
		},
		{
			name:                          "get 1 product when filtering by priceLessThan",
			filterByAmountLessThanOrEqual: 59000,
			expectedProductsResponse:      getExpectedProductResponsesFilteredByPrice(59000),
			expectedNumberOfProducts:      1, // As there are only 1 items in with price less than or equal 59000
		},
		{
			name:                          "get no product when filtering by priceLessThan",
			filterByAmountLessThanOrEqual: 50000,
			expectedProductsResponse:      []types.ProductResponse{},
			expectedNumberOfProducts:      0, // As there are only 1 items in with price less than or equal 59000
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/products", nil)
			// Query params
			params := url.Values{}
			params.Add("category", testCase.filterByCategory)
			params.Add("priceLessThan", fmt.Sprintf("%d", testCase.filterByAmountLessThanOrEqual))
			req.URL.RawQuery = params.Encode()

			require.NoError(t, err)
			router.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
			outputResponse := []types.ProductResponse{}
			err = json.Unmarshal(w.Body.Bytes(), &outputResponse)
			require.NoError(t, err)
			assert.Equal(t, testCase.expectedNumberOfProducts, len(outputResponse))
			assert.Equal(t, testCase.expectedProductsResponse, outputResponse)
		})
	}
}

func getExpectedProductResponseFilteredByCategoryAndPrice(category string, price int) []types.ProductResponse {
	var filteredProductResponses []types.ProductResponse
	allProductResponses := getExpectedProductResponseForAllProducts()
	allProductResponses = append(allProductResponses, types.ProductResponse{
		SKU:      "000006",
		Name:     "Gucci cap",
		Category: "caps",
		ProductPricing: types.ProductPricingModel{
			OriginalPrice: 65000,
			FinalPrice:    65000,
			Currency:      DefaultCurrency,
		},
	})

	for _, productResponse := range allProductResponses {
		if productResponse.ProductPricing.OriginalPrice <= price && productResponse.Category == category {
			filteredProductResponses = append(filteredProductResponses, productResponse)
		}
	}
	return filteredProductResponses
}

func getExpectedProductResponsesFilteredByPrice(price int) []types.ProductResponse {
	var filteredProductResponses []types.ProductResponse
	allProductResponses := getExpectedProductResponseForAllProducts()

	for _, productResponse := range allProductResponses {
		if productResponse.ProductPricing.OriginalPrice <= price {
			filteredProductResponses = append(filteredProductResponses, productResponse)
		}
	}
	return filteredProductResponses
}

func getExpectedProductResponsesFilteredByCategory(category string) []types.ProductResponse {
	var filteredProductResponses []types.ProductResponse
	allProductResponses := getExpectedProductResponseForAllProducts()

	for _, productResponse := range allProductResponses {
		if productResponse.Category == category {
			filteredProductResponses = append(filteredProductResponses, productResponse)
		}
	}
	return filteredProductResponses
}

func getExpectedProductResponseForAllProducts() []types.ProductResponse {
	discount30 := "30%"
	return []types.ProductResponse{
		{
			SKU:      "000001",
			Name:     "BV Lean Leather ankle boots",
			Category: "boots",
			ProductPricing: types.ProductPricingModel{
				OriginalPrice:      89000,
				DiscountPercentage: &discount30,
				FinalPrice:         62300,
				Currency:           DefaultCurrency,
			},
		},
		{
			SKU:      "000002",
			Name:     "BV Lean Leather ankle boots",
			Category: "boots",
			ProductPricing: types.ProductPricingModel{
				OriginalPrice:      99000,
				DiscountPercentage: &discount30,
				FinalPrice:         69300,
				Currency:           DefaultCurrency,
			},
		},
		{
			SKU:      "000003",
			Name:     "Ashlinton leather ankle boots",
			Category: "boots",
			ProductPricing: types.ProductPricingModel{
				OriginalPrice:      71000,
				DiscountPercentage: &discount30,
				FinalPrice:         49700,
				Currency:           DefaultCurrency,
			},
		},
		{
			SKU:      "000004",
			Name:     "Naima embellished suede sandals",
			Category: "sandals",
			ProductPricing: types.ProductPricingModel{
				OriginalPrice: 79500,
				FinalPrice:    79500,
				Currency:      DefaultCurrency,
			},
		},
		{
			SKU:      "000005",
			Name:     "Nathane leather sneakers",
			Category: "sneakers",
			ProductPricing: types.ProductPricingModel{
				OriginalPrice: 59000,
				FinalPrice:    59000,
				Currency:      DefaultCurrency,
			},
		},
	}
}
