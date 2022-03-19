package ancestry

import (
	"hsecode.com/stdlib/tree"
)

type time struct {
	enter, exit int
}
type Ancestry struct {
	h map[int]time
}

func New(T *tree.Tree) *Ancestry {
	a := new(Ancestry)
	a.h = make(map[int]time)
	dfs_timer := 0
	var dfs func(*tree.Tree)
	dfs = func(root *tree.Tree) {
		dfs_timer += 1
		enter := dfs_timer
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
		dfs_timer += 1
		exit := dfs_timer
		a.h[root.Value] = time{enter, exit}
	}
	dfs(T)
	return a
}

func (A *Ancestry) IsDescendant(a, b int) bool {
	v1, ok1 := A.h[a]
	v2, ok2 := A.h[b]
	if !ok1 || !ok2 {
		panic("")
	}
	if a != b {
		if v1.enter < v2.enter && v1.exit > v2.exit {
			return true
		}
	}
	return false
}
