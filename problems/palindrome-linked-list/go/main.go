package main

import (
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {

}

// palindromeSlice uses one slice to compare forward and backward
// time O(n)
// space O(n)
func palindromeSlice(head *ListNode) bool {
	current := head
	slice := make([]int, 0)

	for current != nil {
		slice = append(slice, current.Val)
		current = current.Next
	}

	for i, j := 0, len(slice); i < j; i++ {
		j--
		if slice[i] != slice[j] {
			return false
		}
	}

	return true
}

// palindromeStrs uses strings to compare forward and backward
// time O(n)
// space O(2n)
func palindromeStrs(head *ListNode) bool {
	var forward string
	var backward string
	current := head
	for current != nil {
		value := fmt.Sprint(current.Val)
		forward = forward + value
		backward = value + backward
		current = current.Next
	}
	return forward == backward
}
