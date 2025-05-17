package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"food_ordering_api/internal/utils"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// fetchProductData is a helper function to fetch product data from the given URL
func fetchProductData(url string) ([]byte, error) {
	data, err := utils.FetchData(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from %s: %w", url, err)
	}
	return data, nil
}

// FetchProducts retrieves all products from the external API
func FetchProducts() ([]Product, error) {
	url := "https://orderfoodonline.deno.dev/api/product"
	data, err := fetchProductData(url)
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, errors.New("failed to parse product data")
	}

	return products, nil
}

// FetchProductByID retrieves a single product by ID from the external API
func FetchProductByID(id string) (*Product, error) {
	url := fmt.Sprintf("https://orderfoodonline.deno.dev/api/product/%s", id)
	data, err := fetchProductData(url)
	if err != nil {
		return nil, err
	}

	var product Product
	if err := json.Unmarshal(data, &product); err != nil {
		return nil, errors.New("failed to parse product data")
	}

	return &product, nil
}
