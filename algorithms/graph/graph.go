package graph

type Type int

const (
	Directed Type = iota
	Undirected
)

type Graph struct {
	Type     Type
	Lookup   map[int]Node
	Alledges map[int]map[int]interface{}
}

type Node interface {
	ID() int
}

func New(graphType Type) *Graph {
	g := new(Graph)
	g.Type = graphType
	g.Alledges = make(map[int]map[int]interface{})
	g.Lookup = make(map[int]Node)
	return g
}

func (g *Graph) AddNode(node Node) {
	g.Lookup[node.ID()] = node
}

func (g *Graph) Node(id int) (Node, bool) {
	if g.Lookup[id] == nil {
		return nil, false
	}
	return g.Lookup[id], true
}

func (g *Graph) AddEdge(u, v int, edgeData interface{}) {
	if g.Lookup[u] == nil || g.Lookup[v] == nil {
		panic("fjnfj")
	}
	if g.Alledges[u] == nil {
		g.Alledges[u] = make(map[int]interface{})
	}
	g.Alledges[u][v] = edgeData
	if g.Type == Undirected {
		if g.Alledges[v] == nil {
			g.Alledges[v] = make(map[int]interface{})
		}
		g.Alledges[v][u] = edgeData
	}
}

func (g *Graph) Edge(u, v int) (interface{}, bool) {
	if g.Alledges[u][v] == nil {
		return nil, false
	}
	return g.Alledges[u][v], true

}
func (g *Graph) Edges(f func(u, v Node, edgeData interface{})) {
	if g.Type == Directed {
		for i, r := range g.Alledges {
			for w, _ := range r {
				f(g.Lookup[i], g.Lookup[w], g.Alledges[i][w])
			}
		}
	} else {
		c := make(map[int][]int, 0)
		for i, r := range g.Alledges {
			for w := range r {
				e := 0
				for s := 0; s < len(c[w]); s++ {
					if c[w][s] == i {
						e = 1
					}
				}
				if e != 1 {
					c[i] = append(c[i], w)
					f(g.Lookup[i], g.Lookup[w], g.Alledges[i][w])
				}
			}
		}
	}
}

func (g *Graph) Nodes(f func(Node)) {
	for i, _ := range g.Lookup {
		f(g.Lookup[i])
	}
}

func (g *Graph) Neighbours(u int, f func(v Node, edgeData interface{})) {
	if g.Lookup[u] == nil {
		panic("jdn;k")
	}
	for i, _ := range g.Alledges[u] {
		f(g.Lookup[i], g.Alledges[u][i])
	}
}
