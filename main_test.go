package main

import "testing"

func TestScanItems(t *testing.T) {
	testCases := []struct {
		name      string
		discounts Discounts
		basket    Basket
		products  Products
		expected  float64
	}{
		{
			name: "1. returns the correct price for 3 normal items",
			discounts: Discounts{
				{
					Code:        "FR1",
					Price:       3.11,
					BuyOneFree:  true,
					ApplyAt:     0,
					Description: "Buy one get one free",
				},
				{
					Code:        "SR1",
					Price:       0.50,
					BuyOneFree:  false,
					ApplyAt:     3,
					Description: "Buy 3 or more for a reduced price of £4.50",
				},
			},
			basket: Basket{
				Codes: []string{
					"FR1",
					"SR1",
					"CF1",
				},
			},
			products: Products{
				{
					Code:        "FR1",
					Name:        "Fruit Tea",
					Price:       3.11,
					HasDiscount: true,
				},
				{
					Code:        "SR1",
					Name:        "Strawberries",
					Price:       5.00,
					HasDiscount: true,
				},
				{
					Code:        "CF1",
					Name:        "Coffee",
					Price:       11.23,
					HasDiscount: false,
				},
			},
			expected: 19.34,
		},
		{
			name: "2. returns the correct price for 4 items with the 2 fruit tea discount offer",
			discounts: Discounts{
				{
					Code:        "FR1",
					Price:       3.11,
					BuyOneFree:  true,
					ApplyAt:     0,
					Description: "Buy one get one free",
				},
				{
					Code:        "SR1",
					Price:       0.50,
					BuyOneFree:  false,
					ApplyAt:     3,
					Description: "Buy 3 or more for a reduced price of £4.50",
				},
			},
			basket: Basket{
				Codes: []string{
					"FR1",
					"FR1",
					"SR1",
					"CF1",
				},
			},
			products: Products{
				{
					Code:        "FR1",
					Name:        "Fruit Tea",
					Price:       3.11,
					HasDiscount: true,
				},
				{
					Code:        "SR1",
					Name:        "Strawberries",
					Price:       5.00,
					HasDiscount: true,
				},
				{
					Code:        "CF1",
					Name:        "Coffee",
					Price:       11.23,
					HasDiscount: false,
				},
			},
			expected: 19.34,
		},
		{
			name: "3. returns the correct price for 6 items where strawberries are brought in bulk (3 or more)",
			discounts: Discounts{
				{
					Code:        "FR1",
					Price:       3.11,
					BuyOneFree:  true,
					ApplyAt:     0,
					Description: "Buy one get one free",
				},
				{
					Code:        "SR1",
					Price:       0.50,
					BuyOneFree:  false,
					ApplyAt:     3,
					Description: "Buy 3 or more for a reduced price of £4.50",
				},
			},
			basket: Basket{
				Codes: []string{
					"FR1",
					"SR1",
					"SR1",
					"SR1",
					"SR1",
					"CF1",
				},
			},
			products: Products{
				{
					Code:        "FR1",
					Name:        "Fruit Tea",
					Price:       3.11,
					HasDiscount: true,
				},
				{
					Code:        "SR1",
					Name:        "Strawberries",
					Price:       5.00,
					HasDiscount: true,
				},
				{
					Code:        "CF1",
					Name:        "Coffee",
					Price:       11.23,
					HasDiscount: false,
				},
			},
			expected: 32.34,
		},
		{
			name: "4. returns the correct price for 8 items where strawberries are brought in bulk (3 or more) and fruit tea is an odd amount",
			discounts: Discounts{
				{
					Code:        "FR1",
					Price:       3.11,
					BuyOneFree:  true,
					ApplyAt:     0,
					Description: "Buy one get one free",
				},
				{
					Code:        "SR1",
					Price:       0.50,
					BuyOneFree:  false,
					ApplyAt:     3,
					Description: "Buy 3 or more for a reduced price of £4.50",
				},
			},
			basket: Basket{
				Codes: []string{
					"FR1",
					"FR1",
					"FR1",
					"SR1",
					"SR1",
					"SR1",
					"SR1",
					"CF1",
				},
			},
			products: Products{
				{
					Code:        "FR1",
					Name:        "Fruit Tea",
					Price:       3.11,
					HasDiscount: true,
				},
				{
					Code:        "SR1",
					Name:        "Strawberries",
					Price:       5.00,
					HasDiscount: true,
				},
				{
					Code:        "CF1",
					Name:        "Coffee",
					Price:       11.23,
					HasDiscount: false,
				},
			},
			expected: 35.45,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BasketCost(tc.basket, tc.products, tc.discounts)

			if got != tc.expected {
				t.Errorf("ScanItems(%v, %v) == %v, want %v", tc.basket, tc.products, got, tc.expected)
			}
		})
	}
}
