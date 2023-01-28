package main

import "testing"

func BuildList(numbers []int) *ListNode {
	head := &ListNode{Val: numbers[0]}
	current := head
	for i := 1; i < len(numbers); i++ {
		current.Next = &ListNode{Val: numbers[i]}
		current = current.Next
	}
	return head
}

func TestPalindrome(t *testing.T) {
	type TC struct {
		name     string
		input    *ListNode
		expected bool
	}
	tt := []TC{
		{
			name:     "case1",
			expected: true,
			input:    BuildList([]int{1, 2, 2, 1}),
		},
		{
			name:     "case2",
			expected: false,
			input:    BuildList([]int{1, 2}),
		},
		{
			name:     "case3",
			expected: true,
			input:    BuildList([]int{1, 1, 1}),
		},
		{
			name:     "case4",
			expected: true,
			input:    BuildList([]int{4, 4}),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := palindromeSlice(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v got %v slice", tc.expected, result)
			}

			result = palindromeStrs(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v got %v in strs", tc.expected, result)
			}
		})
	}
}
