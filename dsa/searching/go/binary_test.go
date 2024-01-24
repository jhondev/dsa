package searching

import "testing"

func TestSearchBinary(t *testing.T) {
	// create test cases
	testCases := []struct {
		name     string
		src      []int
		value    int
		expected int
		found    bool
	}{
		{
			name:     "empty slice",
			src:      []int{},
			value:    1,
			expected: -1,
			found:    false,
		},
		{
			name:     "one element slice",
			src:      []int{1},
			value:    1,
			expected: 0,
			found:    true,
		},
		{
			name:     "two element slice (value found)",
			src:      []int{1, 2},
			value:    1,
			expected: 0,
			found:    true,
		},
		{
			name:     "two element slice (value not found)",
			src:      []int{1, 2},
			value:    3,
			expected: -1,
			found:    false,
		},
		{
			name:     "three element slice (value found)",
			src:      []int{1, 2, 3},
			value:    2,
			expected: 1,
			found:    true,
		},
		{
			name:     "three element slice (value not found)",
			src:      []int{1, 2, 3},
			value:    4,
			expected: -1,
			found:    false,
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, found := SearchBinary(tc.src, tc.value)
			if actual != tc.expected || found != tc.found {
				t.Errorf("SearchBinary(%v, %d) = (%d, %t), expected (%d, %t)", tc.src, tc.value, actual, found, tc.expected, tc.found)
			}
		})
	}
}
