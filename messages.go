package moj

import "fmt"

// PrintMessages will print messages to the console for the user
func PrintMessages(products Products, discounts Discounts) {
	printWelcomeMessage()
	printCurrentProducts(products)
	printDiscountProducts(discounts, products)
}

// printWelcomeMessage will print the welcome message to the console
func printWelcomeMessage() {
	welcomeMessage :=
		`
|--------------------------------|
|            Hi there            |
|           Welcome to           |
|     The Local Shopping Mall    |
|--------------------------------|
		`
	fmt.Println(welcomeMessage)
}

// printCurrentProducts will print the current list of products and their prices
func printCurrentProducts(products Products) {
	currentProducts :=
		`
===============================
=       Here is a list of     =
=      our current products   =
===============================
		`
	fmt.Println(currentProducts)
	for _, product := range products {
		fmt.Printf("Product: %v \n", product.Name)
		fmt.Printf("Price: £%.2f \n\n", product.Price)
	}
}

// printDiscountProducts will print the current list of products and their prices
func printDiscountProducts(discounts Discounts, products Products) {
	currentDiscounts :=
		`
===============================
=       Here is a list of     =
=     our current discounts   =
===============================
		`
	fmt.Println(currentDiscounts)
	for _, discount := range discounts {
		product := findProduct(products, discount.Code)
		fmt.Printf("%v Discount offer: %v \n\n", product.Name, discount.Description)
	}
}
