package main

import (
	"fmt"
	"testing"
)

func TestCalculateTurns(t *testing.T) {
	tt := []struct {
		seats     int
		customer  int
		dishes    string
		rotations int
	}{
		{seats: 3, customer: 0, dishes: "EPD", rotations: 2},
		{seats: 7, customer: 0, dishes: "ENPNDNN", rotations: 4},
		{seats: 7, customer: 1, dishes: "ENPDNNN", rotations: 4},
		{seats: 9, customer: 1, dishes: "ENNEDPNNN", rotations: 5},
	}

	for _, tc := range tt {
		t.Run(tc.dishes, func(t *testing.T) {
			rotations := CalculateTurns(tc.seats, tc.customer, tc.dishes)
			if rotations != tc.rotations {
				t.Fatal(fmt.Sprintf("expecting %v got %v", tc.rotations, rotations))
			}
		})
	}
}
