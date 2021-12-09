package pool

import (
	"context"
	"fmt"
)

/***Graph representation***/

func (g *Instance_Graph) AddInstance(n *Instance_Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *Instance_Graph) AddEdge(v1, v2 *Instance_Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[context.Context][]*Instance_Node)
	}
	g.edges[v1.ctx] = append(g.edges[v1.ctx], v2)
	g.edges[v2.ctx] = append(g.edges[v2.ctx], v1)

	g.lock.Unlock()
}

func (g *Instance_Graph) toString() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {

		//print current vertex
		s += g.nodes[i].toString() + "-->"

		//grab neighbors and print them
		neighbors := g.edges[g.nodes[i].ctx]
		for j := 0; j < len(neighbors); j++ {
			s += neighbors[j].toString() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}

/***End of Graph representation***/
