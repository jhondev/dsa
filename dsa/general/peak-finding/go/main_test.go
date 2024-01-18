package main

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestFindPeak(t *testing.T) {
	tt := []struct {
		data     []int
		expected map[int]bool
	}{
		{
			data:     []int{5, 10, 18, 15, 14, 11, 10, 9, 4},
			expected: map[int]bool{18: true},
		},
		{
			data:     []int{5, 10, 20, 15},
			expected: map[int]bool{20: true},
		},
		{
			data:     []int{10, 20, 15, 2, 23, 90, 67},
			expected: map[int]bool{20: true, 90: true},
		},
		{
			data:     []int{134},
			expected: map[int]bool{134: true},
		},
		{
			data:     []int{10, 15, 3, 4},
			expected: map[int]bool{15: true, 4: true},
		},
		{
			data:     []int{10, 67},
			expected: map[int]bool{67: true},
		},
	}
	for _, tc := range tt {
		peak := FindPeak(tc.data)
		if !tc.expected[peak] {
			t.Errorf("\nresult should be one of %v got %d", maps.Keys(tc.expected), peak)
		}
	}
}
