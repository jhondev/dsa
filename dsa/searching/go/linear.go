package searching

func SearchLinear(src []int, value int) (int, bool) {
	for i, v := range src {
		if v == value {
			return i, true
		}
	}
	return -1, false
}
