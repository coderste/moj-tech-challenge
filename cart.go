package moj

import (
	"flag"
	"fmt"
)

// Basket is a representation of a user's checkout items
type Basket struct {
	ItemCount map[string]int
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

	// Loop over each product in the basket
	for code, quantity := range basket.ItemCount {
		product := findProduct(products, code)
		totalCost += product.Price * float64(quantity)

		// Check if the product has a discount to apply
		if product.HasDiscount {
			discount := findDiscount(discounts, product.Code)

			// If the product is a buy one get one free discount?
			if discount.BuyOneFree {
				// If there is an even amount of products then the calculation is easy
				// we just cut the amount of items in half and then times that by
				// the product price. If there is an odd number of items that
				// is equal to 3 or more then we need to take one out of the quantity
				// and then cut that quantity in half to get the saving cost
				if isEven(quantity) {
					offerSaving := quantity / 2
					savingCost += product.Price * float64(offerSaving)
				} else if !isEven(quantity) && quantity >= 3 {
					offerSaving := (quantity - 1) / 2
					savingCost += product.Price * float64(offerSaving)
				}
			} else {
				// If the product isn't a buy one get one free
				// that means it's a bulk discount and all we do for that
				// is times the discount price by the amount of products for that discount
				if quantity >= discount.ApplyAt {
					savingCost += discount.Price * float64(quantity)
				}
			}
		}
	}

	return totalCost - savingCost
}

// isEven will check if an item count is even
func isEven(number int) bool {
	return number%2 == 0
}
