package main

import "fmt"

func main() {
	peak := FindPeak([]int{10, 15, 3, 4})
	fmt.Printf("Peak: %d", peak)
}

// TODO: improve algorithm decrease conditionals
func FindPeak(data []int) int {
	n := len(data)
	if n == 0 {
		panic("data should contains numbers")
	}
	if n == 1 {
		return data[0]
	}
	if n == 2 {
		if data[0] > data[1] {
			return data[0]
		} else {
			return data[1]
		}
	}
	// use divide and conquer with binary search
	mindex := n / 2         // middle index
	left := data[mindex-1]  // left number
	mid := data[mindex]     // middle number
	right := data[mindex+1] // right number

	if mid > left && mid > right {
		return mid
	}
	if left > mid {
		return FindPeak(data[:mindex+1])
	}
	if right > mid {
		return FindPeak(data[mindex:])
	}

	return mid
}
