package searching

func SearchBinary(src []int, value int) (int, bool) {
	if len(src) == 0 {
		return -1, false
	}
	start := 0
	end := len(src)
	for start < end {
		mid := start + (end-start)/2
		if src[mid] == value {
			return mid, true
		} else if src[mid] < value {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1, false
}
