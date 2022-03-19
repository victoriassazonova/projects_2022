package tree

import (
	"sort"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func NewBST(elements []int) *Tree {
	unique := make([]int, 0)
	h := make(map[int]struct{})
	for _, e := range elements {
		_, ok := h[e]
		if ok {
			continue
		}
		h[e] = struct{}{}
		unique = append(unique, e)
	}
	sort.Ints(unique)
	return ArrayToBinaryTree(unique, 0, len(unique)-1)
}

func ArrayToBinaryTree(a []int, start int, end int) *Tree {
	if start > end {
		var returnTree *Tree
		return returnTree
	}
	middle := (start + end) / 2
	tree := Tree{Value: a[middle]}
	tree.Left = ArrayToBinaryTree(a, start, middle-1)
	tree.Right = ArrayToBinaryTree(a, middle+1, end)
	return &tree
}

func (T *Tree) NoLeft() *Tree {
	if T == nil {
		return T
	}
	if T.Left == nil && T.Right == nil {
		return T
	}
	if T.Left == nil {
		T.Right = T.Right.NoLeft()
		return T
	}
	T.Right = T.Right.NoLeft()
	l := T.Left.NoLeft()
	t := l
	for t.Right != nil {
		t = t.Right
	}
	t.Right = T
	T.Left = nil
	return l
}
