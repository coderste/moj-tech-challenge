package moj

import (
	"flag"
	"fmt"
)

// Basket is a representation of a user's checkout items
type Basket struct {
	Codes []string
}

// CartCLI will initiate the command line interface of the application
func CartCLI() {
	// Setting up application flags
	productFlag := flag.String("products", "files/products.csv", "location of the product list")
	discountFlag := flag.String("discounts", "files/discounts.csv", "location of the discount list")
	flag.Parse()

	// Load in the products and discounts
	products := LoadProducts(productFlag)
	discounts := LoadDiscounts(discountFlag)

	// check if a product has a discount
	products = CheckProductDiscount(products, discounts)

	// Print the CLI messages
	PrintMessages(products, discounts)

	// Items inside the users basket
	basket := UserInput(products)

	// Work out the total cost of the basket
	cost := BasketCost(basket, products, discounts)
	fmt.Printf("Your total basket cost is: £%.2f\n", cost)
}

// BasketCost takes in a list of items in the users basket and will loop through
// each item and add the item price to a total cost
func BasketCost(basket Basket, products Products, discounts Discounts) float64 {
	var totalCost float64
	var savingCost float64

	for _, item := range basket.Codes {
		product := findProduct(products, item)
		totalCost += product.Price
	}

	// discount values
	for _, product := range products {
		if contains(basket, product) {
			if product.HasDiscount {
				discount := findDiscount(discounts, product.Code)
				itemCount := itemCount(basket, product.Code)

				if discount.BuyOneFree {
					if isEven(itemCount) {
						offer := itemCount / 2
						savingCost += product.Price * float64(offer)
					} else if !isEven(itemCount) && itemCount >= 3 {
						offer := (itemCount - 1) / 2
						savingCost += product.Price * float64(offer)
					}
				} else {
					if itemCount >= discount.ApplyAt {
						savingCost += discount.Price * float64(itemCount)
					}
				}
			}
		}
	}

	return totalCost - savingCost
}

// itemCount will count the amount of times
// an item appears in the basket
func itemCount(basket Basket, code string) int {
	var counter int
	for _, item := range basket.Codes {
		if item == code {
			counter++
		}
	}

	return counter
}

// isEven will check if an item count is even
func isEven(number int) bool {
	return number%2 == 0
}

// contains will check the user's basket contains a product's code
func contains(basket Basket, product Product) bool {
	for _, item := range basket.Codes {
		if item == product.Code {
			return true
		}
	}

	return false
}
