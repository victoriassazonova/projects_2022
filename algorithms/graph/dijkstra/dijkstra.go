package dijkstra

import (
	"container/heap"
	"hsecode.com/stdlib/graph"
)

type distanceNode struct {
	node graph.Node
	dist int
}

type priorityQueue []distanceNode

func (q priorityQueue) Len() int            { return len(q) }
func (q priorityQueue) Less(i, j int) bool  { return q[i].dist < q[j].dist }
func (q priorityQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *priorityQueue) Push(n interface{}) { *q = append(*q, n.(distanceNode)) }
func (q *priorityQueue) Pop() interface{} {
	t := *q
	var n interface{}
	n, *q = t[len(t)-1], t[:len(t)-1]
	return n
}

type Path struct {
	Nodes  []graph.Node
	Weight uint
}

const (
	MaxInt64 = 1<<63 - 1
)

func New(g *graph.Graph, uid, vid int, weight func(interface{}) uint) *Path {
	if g.Lookup[uid] == nil || g.Lookup[vid] == nil {
		return nil
	}
	//visited := make(map[int]bool)
	weights := make(map[int]int)
	prev := make(map[graph.Node]int)
	for i := 0; i < len(g.Lookup); i++ {
		weights[i] = MaxInt64
	}
	weights[uid] = 0
	Q := priorityQueue{{node: g.Lookup[uid], dist: 0}}
	for Q.Len() != 0 {
		mid := heap.Pop(&Q).(distanceNode)
		node := mid.node.ID()
		if node == vid {
			p := new(Path)
			p.Weight = uint(weights[vid])
			u := vid
			if prev[g.Lookup[u]] != 0 {
				for u != uid {
					p.Nodes = append(p.Nodes, g.Lookup[u])
					u = prev[g.Lookup[u]]
				}
			}
			p.Nodes = append(p.Nodes, g.Lookup[uid])
			for i, j := 0, len(p.Nodes)-1; i < j; i, j = i+1, j-1 {
				p.Nodes[i], p.Nodes[j] = p.Nodes[j], p.Nodes[i]
			}
			return p
		}
		//min := MaxInt64
		//node := 0
		//for t, j := range weights {
		//	if j < min && visited[t] != true {
		//		min = j
		//		node = t
		//	}
		//}
		//visited[node] = true
		//fmt.Println(node)
		for i := range g.Alledges[node] {
			if weights[i] > weights[node]+int(weight(g.Alledges[node][i])) {
				weights[i] = weights[node] + int(weight(g.Alledges[node][i]))
				prev[g.Lookup[i]] = node
				heap.Push(&Q, distanceNode{node: g.Lookup[i], dist: weights[i]})
				//if min>weights[i]{
				//	min = weights[i]
				//	node = i
				//}
				//fmt.Println(weights[i])
			}
		}
		//fmt.Println(weights)

	}
	//fmt.Println(weights[vid])
	return nil
}
