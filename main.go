package main

import (
	"fmt"
	"math"
)

// Challenge: Implement a checkout so that it will scan items in and calculate total
// prices correctly for any combination of products and offers above

// Shopping Items:
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

type Items []Item

type Item struct {
	Code  string
	Name  string
	Price float64
}

var BaseItems = Items{
	{
		Code:  "FR1",
		Name:  "Fruit Tea",
		Price: 3.11,
	},
	{
		Code:  "SR1",
		Name:  "Strawberries",
		Price: 5.00,
	},
	{
		Code:  "CF1",
		Name:  "Coffee",
		Price: 11.23,
	},
}
var ItemCosts = map[string]float64{
	"FR1": 3.11,
	"SR1": 5.00,
	"CF1": 11.23,
}

func main() {
	price := ScanItems(BaseItems)
	fmt.Println(price)
}

// ScanItems takes in a list of items and will loop through
// each item and add the item price to a total cost
func ScanItems(items Items) float64 {
	var totalCost float64
	var fruitTeaCount int
	var strawberryCount int

	for _, item := range items {
		totalCost += item.Price

		if item.Code == "FR1" {
			fruitTeaCount++
		}

		if item.Code == "SR1" {
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
