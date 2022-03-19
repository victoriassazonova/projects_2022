package tsort

import (
	"errors"
	"fmt"
	"hsecode.com/stdlib/graph"
)

//func New(g *graph.Graph) ([]graph.Node, error) {
//	if g.Type == graph.Undirected {
//		return []graph.Node{}, errors.New("wrong1")
//	}
//	var L []graph.Node
//	L = make([]graph.Node, 0)
//	visited := make(map[graph.Node]int)
//	var visit func(graph.Node) error
//	visit = func(node graph.Node) error {
//		stack := make([]graph.Node, 0)
//		stack = append(stack, node)
//		for len(stack) != 0 {
//			s := stack[0]
//			stack = stack[1:]
//			if visited[s] == 2 {
//				return errors.New("wrong1")
//			}
//			visited[s] = 1
//			L = append(L, s)
//
//			for i := range g.Alledges[s.ID()] {
//				if visited[g.Lookup[i]] == 0 {
//					stack = append(stack, g.Lookup[i])
//				}
//				if visited[g.Lookup[i]] == 1 {
//					return errors.New("wrong2")
//				}
//			}
//			visited[s] = 2
//		}
//		return nil
//	}
//	for i, _ := range g.Lookup {
//		if visited[g.Lookup[i]] == 0 {
//			e := visit(g.Lookup[i])
//			if e != nil {
//				return []graph.Node{}, e
//			}
//		}
//	}
//	for i, j := 0, len(L)-1; i < j; i, j = i+1, j-1 {
//		L[i], L[j] = L[j], L[i]
//	}
//	return L, nil
//}

func New(g *graph.Graph) ([]graph.Node, error) {
	if g.Type == graph.Undirected {
		return []graph.Node{}, errors.New("wrong1")
	}
	L := make([]graph.Node, 0)
	nodes := make(map[int]bool)
	for i, _ := range g.Lookup {
		nodes[i] = true
	}
	for len(nodes) > 0 {
		ch := 0
		for value := range nodes {
			t := 0
			for node := range g.Alledges[value] {
				if nodes[node] == true {
					t = 1
				}
			}
			if t == 0 {
				L = append(L, g.Lookup[value])
				delete(nodes, value)
				ch = 1
			}
		}
		if ch == 0 {
			return nil, fmt.Errorf("wrong2")
		}
	}
	for i, j := 0, len(L)-1; i < j; i, j = i+1, j-1 {
		L[i], L[j] = L[j], L[i]
	}
	return L, nil
}
