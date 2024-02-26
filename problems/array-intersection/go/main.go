package main

import (
	"slices"
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
		for range rep {
			total = append(total, k)
		}
	}

	slices.Sort(total)
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
