package moj

import (
	"errors"
	"fmt"
)

var (
	tooManyItems   = errors.New("You have selected too many items. A maximum of 100 items allowed per product")
	notEnoughItems = errors.New("Negative item amounts are not allowed. Please select a number of 0 or greater")
)

// UserInput will watch the console for the users input and based on how many
// items they select for each product will be enter
func UserInput(products Products) Basket {
	var basket Basket
	basket.ItemCount = make(map[string]int)

	for _, product := range products {
		var response int

		for {
			fmt.Printf("How many %v would you like to buy? ", product.Name)
			_, err := fmt.Scan(&response)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if ok, err := validResponse(response); ok {
				break
			} else {
				fmt.Println(err)
				continue
			}
		}

		basket.ItemCount[product.Code] = response
	}

	return basket
}

func validResponse(response int) (bool, error) {
	if response > 100 {
		return false, tooManyItems
	}

	if response < 0 {
		return false, notEnoughItems
	}

	return true, nil
}
