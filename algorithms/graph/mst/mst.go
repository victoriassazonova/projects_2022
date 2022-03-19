package mst

import (
	"hsecode.com/stdlib/graph"
)

const (
	MaxInt64 = 1<<63 - 1
)

func New(g *graph.Graph, weight func(interface{}) int) *graph.Graph {
	if g.Type == graph.Directed {
		panic("directed")
	}
	if len(g.Lookup) < 2 {
		return g
	}
	L := graph.New(graph.Undirected)
	v := make([]int, 0)
	k := 0
	for u := range g.Lookup {
		if k == 0 {
			v = append(v, u)
			L.AddNode(g.Lookup[u])
			k = 1
		}
		break
	}
	for len(L.Lookup) != len(g.Lookup) {
		min := MaxInt64
		var toadd []int
		for s := 0; s < len(v); s++ {
			for i := range g.Alledges[v[s]] {
				if weight(g.Alledges[v[s]][i]) < min && L.Lookup[i] == nil && g.Lookup[i] != nil {
					min = weight(g.Alledges[v[s]][i])
					toadd = []int{v[s], i}
				}
			}
		}
		if len(toadd) == 2 {
			L.AddNode(g.Lookup[toadd[1]])
			v = append(v, toadd[1])
			L.AddEdge(toadd[0], toadd[1], g.Alledges[toadd[0]][toadd[1]])
			delete(g.Alledges[toadd[0]], toadd[1])
		} else {
			k := 0
			for u := range g.Lookup {
				if k == 0 && L.Lookup[u] == nil {
					v = append(v, u)
					L.AddNode(g.Lookup[u])
					k = 1
				}
				break
			}
		}

	}
	return L
}
