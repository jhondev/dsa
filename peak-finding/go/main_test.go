package main

import "testing"

func TestFindPeak(t *testing.T) {
	tt := []struct {
		data     []int
		expected int
	}{
		{
			data:     []int{5, 10, 20, 15, 14, 11, 10, 9, 4},
			expected: 20,
		},
		{
			data:     []int{5, 10, 20, 15},
			expected: 20,
		},
	}
	for _, tc := range tt {
		peak := FindPeak(tc.data)
		if peak != 20 {
			t.Errorf("\nresult should be %d got %d", tc.expected, peak)
		}
	}
}
