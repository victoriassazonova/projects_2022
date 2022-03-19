package xsort

import (
	"container/list"
	"math"
)

//func Sort(data *list.List, less func(a, b *list.Element) bool) {
//	current_size := 1
//	for current_size < data.Len()-1 {
//		left := 0
//		for left < data.Len()-1 {
//			mid := math.Min(float64(left+current_size-1), float64(data.Len()-1))
//			right := math.Min(float64(left+2*current_size-1), float64(data.Len()-1))
//			merge(data, left, int(mid), int(right))
//			left = left + current_size*2
//		}
//		current_size = 2 * current_size
//	}
//}
