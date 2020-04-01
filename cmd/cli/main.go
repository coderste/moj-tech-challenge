package main

import "github.com/moretonb/moj-coderste-tech-challenge"

// Challenge: Implement a checkout so that it will scan items in and calculate total
// prices correctly for any combination of products and offers above

// Shopping Discounts:
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

func main() {
	moj.CartCLI()
}
