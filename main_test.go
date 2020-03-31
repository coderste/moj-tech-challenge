package main

import "testing"

func TestScanItems(t *testing.T) {
	testCases := []struct {
		name     string
		items    Items
		expected float64
	}{
		{
			name: "1. returns the correct price for 3 normal items",
			items: Items{
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
			},
			expected: 19.34,
		},
		{
			name: "2. returns the correct price for 4 items with the 2 fruit tea discount offer",
			items: Items{
				{
					Code:  "FR1",
					Name:  "Fruit Tea",
					Price: 3.11,
				},
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
			},
			expected: 19.34,
		},
		{
			name: "3. returns the correct price for 6 items where strawberries are brought in bulk (3 or more)",
			items: Items{
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
					Code:  "SR1",
					Name:  "Strawberries",
					Price: 5.00,
				},
				{
					Code:  "SR1",
					Name:  "Strawberries",
					Price: 5.00,
				},
				{
					Code:  "SR1",
					Name:  "Strawberries",
					Price: 5.00,
				},
			},
			expected: 21.11,
		},
		{
			name: "4. returns the correct price for 8 items where strawberries are brought in bulk (3 or more) and fruit tea is an odd amount",
			items: Items{
				{
					Code:  "FR1",
					Name:  "Fruit Tea",
					Price: 3.11,
				},
				{
					Code:  "FR1",
					Name:  "Fruit Tea",
					Price: 3.11,
				},
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
					Code:  "SR1",
					Name:  "Strawberries",
					Price: 5.00,
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
				{
					Code:  "CF1",
					Name:  "Coffee",
					Price: 11.23,
				},
			},
			expected: 42.18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ScanItems(tc.items)

			if got != tc.expected {
				t.Errorf("ScanItems(%v) == %v, want %v", tc.items, got, tc.expected)
			}
		})
	}
}
