package graph

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Item generic.Type

type Node struct {
	Value Item
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
	lock  sync.RWMutex
}

func CreateGraph() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
		Edges: make(map[Node][]*Node, 0),
		lock:  sync.RWMutex{},
	}
}

// AddVertex adds a vertex.
func (g *Graph) AddVertex(value int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.Nodes = append(g.Nodes, &Node{Value:value})
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node, 0)
	}
	node1 := Node{Value:v1}
	node2 := Node{Value:v2}
	g.Edges[node1] = append(g.Edges[node1], &node2)
	g.Edges[node2] = append(g.Edges[node2], &node1)
}

