package pick

import (
	"container/heap"
)

type Ordered interface {
	Len() int
	Less(i, j int) bool
}

type Heap struct {
	indices []int
	data    Ordered
}

func (h *Heap) Len() int           { return len(h.indices) }
func (h *Heap) Swap(i, j int)      { h.indices[i], h.indices[j] = h.indices[j], h.indices[i] }
func (h *Heap) Less(i, j int) bool { return !h.data.Less(h.indices[i], h.indices[j]) }
func (h *Heap) Push(x interface{}) { h.indices = append(h.indices, x.(int)) }
func (h *Heap) Pop() interface{} {
	old := h.indices
	n := len(old)
	x := old[n-1]
	h.indices = old[0 : n-1]
	return x
}

func FirstN(data Ordered, n int) []int {
	h := Heap{
		indices: make([]int, 0),
		data:    data,
	}
	indices := make([]int, 0)
	if data.Len() > n {
		for i := 0; i < n; i++ {
			heap.Push(&h, i)
		}
		for i := n; i < data.Len(); i++ {
			if !data.Less(h.indices[0], i) {
				heap.Push(&h, i)
				heap.Pop(&h)
			}
		}
		return h.indices
	}
	for i := 0; i < data.Len(); i++ {
		indices = append(indices, i)
	}
	return indices
}
