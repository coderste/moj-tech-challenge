package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Challenge: Implement a checkout so that it will scan items in and calculate total
// prices correctly for any combination of products and offers above

// Shopping Products:
// Code | Name         | Price
// FR1  | Fruit Tea    | £3.11
// SR1  | Strawberries | £5.00
// CF1  | Coffee       | £11.23

// Considerations:
// Fruit tea has a buy one get one free offer so if they have multiple of 2 then we discount
// one as a free item. If it's not divisible by 2 then we charge them for the price of 2 items
// with one still being a free item.
//
// Strawberries has a bulk price option where if you buy 3 or more then you get them for £4.50 instead of £5.00 each
//
// Because the CEO and COO may change their minds the pricing scheme has to be flexible

type Products []Product

type Discounts []Discount

type Discount struct {
	Code        string
	Price       float64
	BuyOneFree  bool
	ApplyAt     int
	Description string
}

type Product struct {
	Code        string
	Name        string
	Price       float64
	HasDiscount bool
}

type Basket struct {
	Codes []string
}

func main() {
	// Command line flags
	productFlag := flag.String("products", "files/products.csv", "location of the product list")
	discountFlag := flag.String("discounts", "files/discounts.csv", "location of the discount list")
	flag.Parse()

	// Open each file
	productsFile, productsFileErr := os.Open(*productFlag)
	if productsFileErr != nil {
		panic(productsFileErr)
	}
	defer productsFile.Close()

	discountsFile, discountsFileErr := os.Open(*discountFlag)
	if discountsFileErr != nil {
		panic(discountsFileErr)
	}
	defer discountsFile.Close()

	// CSV Files for products and discounts
	csvProducts := csv.NewReader(productsFile)
	csvDiscounts := csv.NewReader(discountsFile)

	products := loadProducts(csvProducts)
	discounts := loadDiscounts(csvDiscounts)

	// check if a product has a discount
	products = checkDiscount(products, discounts)

	printMessages(products, discounts)
	basket := userInput(products, discounts)

	cost := BasketCost(basket, products, discounts)
	fmt.Printf("Your total basket cost is: £%.2f\n", cost)
}

func checkDiscount(products Products, discounts Discounts) Products {
	for _, discount := range discounts {
		for i, product := range products {
			if product.Code == discount.Code {
				products[i].HasDiscount = true
			}
		}
	}

	return products
}

func userInput(products Products, discounts Discounts) Basket {
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

func printMessages(products Products, discounts Discounts) {
	welcomeMessage :=
		`
|--------------------------------|
|            Hi there            |
|           Welcome to           |
|     The Local Shopping Mall    |
|--------------------------------|
		`
	fmt.Println(welcomeMessage)

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

func findDiscount(discounts Discounts, code string) Discount {
	for _, discount := range discounts {
		if discount.Code == code {
			return discount
		}
	}

	return Discount{}
}

func loadProducts(file *csv.Reader) Products {
	var products Products
	for {
		var product Product
		record, err := file.Read()
		if err == io.EOF {
			break
		}

		product.Code = record[0]
		product.Name = record[1]
		product.Price, _ = strconv.ParseFloat(record[2], 64)

		products = append(products, product)
	}

	return products
}

func loadDiscounts(file *csv.Reader) Discounts {
	var discounts Discounts
	for {
		var discount Discount
		record, err := file.Read()
		if err == io.EOF {
			break
		}

		discount.Code = record[0]
		discount.Price, _ = strconv.ParseFloat(record[1], 64)
		discount.BuyOneFree, _ = strconv.ParseBool(record[2])
		discount.ApplyAt, _ = strconv.Atoi(record[3])
		discount.Description = record[4]

		discounts = append(discounts, discount)
	}

	return discounts
}

// BasketCost takes in a list of items and will loop through
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

func itemCount(basket Basket, code string) int {
	var counter int
	for _, item := range basket.Codes {
		if item == code {
			counter++
		}
	}

	return counter
}

func contains(basket Basket, product Product) bool {
	for _, item := range basket.Codes {
		if item == product.Code {
			return true
		}
	}

	return false
}

func isEven(number int) bool {
	return number%2 == 0
}
