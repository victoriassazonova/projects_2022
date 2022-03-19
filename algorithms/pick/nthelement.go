package pick

import (
	"math/rand"
	"sort"
)

func partition(data sort.Interface, first int, last int, pivotIndex int) int {
	data.Swap(first, pivotIndex)
	left := first + 1
	right := last
	for left <= right {
		for left <= last && data.Less(left, first) {
			left++
		}
		for right >= first && data.Less(first, right) {
			right--
		}
		if left <= right {
			data.Swap(left, right)
			left++
			right--
		}
	}
	data.Swap(first, right)
	return right
}

func NthElement(data sort.Interface, nth int) {
	first := 0
	last := data.Len() - 1
	for {
		q := partition(data, first, last, rand.Intn(last-first+1)+first)
		if q < nth {
			first = q + 1
		} else if q > nth {
			last = q - 1
		} else {
			break
		}
	}
}
