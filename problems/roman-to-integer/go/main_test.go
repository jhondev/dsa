package main

import "testing"

func TestRomanToInt(t *testing.T) {
	type testcase = struct {
		name     string
		input    string
		expected int
	}
	tt := []testcase{
		{
			name:     "case1",
			input:    "III",
			expected: 3,
		},
		{
			name:     "case2",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "case3",
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := romanToInt(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if result != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, result)
			}
		})
	}
}
