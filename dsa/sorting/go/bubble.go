package bubble

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Ordered](src []T) {
	end := len(src) - 1
	for end > 1 {
		for i := 0; i < end; i++ {
			if src[i] > src[i+1] {
				nxt := src[i+1]
				src[i+1] = src[i]
				src[i] = nxt
			}
		}
		end -= 1
	}
}
