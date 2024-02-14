package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	mapa := summary(a)
	mapb := summary(b)

	total := []int{}
	for k, v := range mapa {
		vb, ok := mapb[k]
		if !ok {
			continue
		}
		rep := 0
		if v <= vb {
			rep = v
		} else {
			rep = vb
		}
		for i := 0; i < rep; i++ {
			total = append(total, k)
		}
	}

	return total
}

func summary(s []int) map[int]int {
	m := make(map[int]int)
	for _, i := range s {
		_, ok := m[i]
		if ok {
			m[i]++
		} else {
			m[i] = 1
		}
	}
	return m
}

func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	fmt.Printf("%v\n", intersection(a, b))
	// output: [2, 23]

	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	fmt.Printf("%v\n", intersection(a, b))
	// output: [1, 1, 1]
}
