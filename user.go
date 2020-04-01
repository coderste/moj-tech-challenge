package moj

import "fmt"

// UserInput will watch the console for the users input and based on how many
// items they select for each product will be enter
func UserInput(products Products) Basket {
	var basket Basket
	for _, product := range products {
		var response int

		fmt.Printf("How many %v would you like to buy? ", product.Name)
		fmt.Scanln(&response)

		for i := 1; i <= response; i++ {
			basket.Codes = append(basket.Codes, product.Code)
		}
	}

	return basket
}
