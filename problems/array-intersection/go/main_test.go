package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	tt := []struct {
		name                  string
		left, right, expected []int
	}{
		{name: "case1", left: []int{23, 3, 1, 2}, right: []int{6, 2, 4, 23}, expected: []int{2, 23}},
		{name: "case2", left: []int{1, 1, 1}, right: []int{1, 1, 1, 1}, expected: []int{1, 1, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := intersection(tc.left, tc.right)
			assert.Equal(t, tc.expected, result)
		})
	}
}
