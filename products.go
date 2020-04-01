package moj

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Products represents a list of Product
type Products []Product

// Product represents a product
type Product struct {
	Code        string
	Name        string
	Price       float64
	HasDiscount bool
}

// LoadProducts takes in the product file location flag
// and will build an list of products to be returned
func LoadProducts(flag *string) Products {
	var products Products

	productsFile, err := os.Open(*flag)
	if err != nil {
		panic(err)
	}
	defer productsFile.Close()

	productCSV := csv.NewReader(productsFile)

	for {
		var product Product
		record, err := productCSV.Read()
		if err == io.EOF {
			break
		}

		// Build a product from the records stored
		// in the CSV
		product.Code = record[0]
		product.Name = record[1]
		product.Price, _ = strconv.ParseFloat(record[2], 64)

		products = append(products, product)
	}

	return products
}

// CheckProductDiscount will loop through each discount and then each product to
// check if that product code is within the discount list.
func CheckProductDiscount(products Products, discounts Discounts) Products {
	for _, discount := range discounts {
		for i, product := range products {
			if product.Code == discount.Code {
				products[i].HasDiscount = true
			}
		}
	}

	return products
}

// findProduct will take a list of products in and a product code
// and return that single item product
func findProduct(products Products, code string) Product {
	for _, product := range products {
		if product.Code == code {
			return product
		}
	}

	return Product{}
}
