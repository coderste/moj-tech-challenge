package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
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
	Code  string
	Name  string
	Price float64
}

type Basket struct {
	Codes []string
}

var ItemCosts = map[string]float64{
	"FR1": 3.11,
	"SR1": 5.00,
	"CF1": 11.23,
}

func main() {
	// Command line flags
	productFlag := flag.String("product file", "products.csv", "file containing a list of products")
	discountFlag := flag.String("discount file", "discounts.csv", "file containing a list of discounted products")
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

	printMessages(products, discounts)
	basket := userInput(products, discounts)

	cost := ScanItems(basket, products)
	fmt.Printf("Your total basket cost is: £%.2f\n", cost)
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

// ScanItems takes in a list of items and will loop through
// each item and add the item price to a total cost
func ScanItems(basket Basket, products Products) float64 {
	var totalCost float64
	var fruitTeaCount int
	var strawberryCount int

	for _, code := range basket.Codes {
		product := findProduct(products, code)
		totalCost += product.Price

		if product.Code == "FR1" {
			fruitTeaCount++
		}

		if product.Code == "SR1" {
			strawberryCount++
		}
	}

	// round the value down to 2 decimal places
	totalCost = math.Floor(totalCost*100) / 100

	if isEven(fruitTeaCount) {
		// 6 fruit teas mean 3 are free so
		// the calculation would look something like
		// 6/2 = 3 totalCost - (ItemCosts * 3 = 9.33)
		fruitTeaCountOffer := fruitTeaCount / 2
		discount := ItemCosts["FR1"] * float64(fruitTeaCountOffer)

		totalCost = totalCost - discount
	} else if !isEven(fruitTeaCount) && fruitTeaCount >= 3 {
		// 3 fruit teas mean 1 is free
		// so we pay for 2 and get 1 free in this case
		// calculation:
		// 3 - 1 = 2 / 2 = 1 = totalCost - (ItemCosts * 1) + 1
		// 3.11 + 3.11 - 3.11 = 6.22
		fruitTeaCountOffer := (fruitTeaCount - 1) / 2
		discount := ItemCosts["FR1"] * float64(fruitTeaCountOffer)

		totalCost = totalCost - discount
	}

	if strawberryCount >= 3 {
		discount := 0.5 * float64(strawberryCount)

		totalCost = totalCost - discount
	}

	return totalCost
}

func isEven(number int) bool {
	return number%2 == 0
}
