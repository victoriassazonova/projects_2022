package tree

import (
	"errors"
	"fmt"
	"strconv"
)

func Decode(data []string) (*Tree, error) {
	d := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		if data[i] != "nil" {
			if n, err := strconv.Atoi(data[i]); err == nil {
				d[i] = n
			} else {
				return nil, errors.New("wrong")
			}
		} else {
			d[i] = -1
			idx_l := i*2 + 1
			if idx_l < len(data) {
				if data[idx_l] != "nil" {
					return nil, errors.New("wrong")
				}
			}
		}
	}
	return dec(d, 0), nil

}
func dec(data []int, idx int) *Tree {
	if len(data) <= idx {
		return nil
	}
	if data[idx] != -1 {
		node := Tree{Value: data[idx]}
		idx_l := idx*2 + 1
		node.Left = dec(data, idx_l)
		node.Right = dec(data, idx*2+2)
		return &node
	}
	return nil
}

func (T *Tree) Encode() []string {
	return levelOrder(T)
}

func levelOrder(root *Tree) []string {
	res := []string{"nil"}
	var dfs func(*Tree, int, int, int)
	dfs = func(root *Tree, level int, s int, q int) {
		if root == nil {
			return
		}
		if s >= len(res) {
			for i := 0; i < q; i++ {
				res = append(res, "nil")
			}
		}

		res[s] = fmt.Sprintf("%d", root.Value)
		dfs(root.Left, level+1, s*2+1, q*2)
		dfs(root.Right, level+1, s*2+2, q*2)
	}

	dfs(root, 0, 0, 1)
	if res[0] == "nil" {
		return []string{}
	}
	for res[len(res)-1] == "nil" {
		res = res[:len(res)-1]
	}

	return res
}
