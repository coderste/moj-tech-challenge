package moj

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Discounts []Discount

type Discount struct {
	Code        string
	Price       float64
	BuyOneFree  bool
	ApplyAt     int
	Description string
}

// LoadProducts takes in the discount file location flag
// and will build an list of discounts to be returned
func LoadDiscounts(flag *string) Discounts {
	var discounts Discounts

	discountsFile, err := os.Open(*flag)
	if err != nil {
		panic(err)
	}
	defer discountsFile.Close()

	discountCSV := csv.NewReader(discountsFile)

	for {
		var discount Discount
		record, err := discountCSV.Read()
		if err == io.EOF {
			break
		}

		// Build a discount from the records stored
		// in the CSV
		discount.Code = record[0]
		discount.Price, _ = strconv.ParseFloat(record[1], 64)
		discount.BuyOneFree, _ = strconv.ParseBool(record[2])
		discount.ApplyAt, _ = strconv.Atoi(record[3])
		discount.Description = record[4]

		discounts = append(discounts, discount)
	}

	return discounts
}

// findDiscount will take a list of products in and a product code
// and return that single item product
func findDiscount(discounts Discounts, code string) Discount {
	for _, discount := range discounts {
		if discount.Code == code {
			return discount
		}
	}

	return Discount{}
}
