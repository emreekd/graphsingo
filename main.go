package main

import (
	graph "./business/graph"
	node "./business/node"
)

func main() {
	// Init the root node
	nodes := &[]node.Node{
		node.Node{
			Value: 1,
			Adjacents: []node.Node{
				node.Node{
					Value: 2,
				},
			},
		},
	}
	graph := &graph.DFSGraph{  //init the graph
		Nodes: nodes,
	}
	// add new adjacents
	graph.Add(1, 3)
	graph.Add(2, 4)
	graph.Add(3, 6)
	graph.Add(3, 7)
	graph.Add(2, 8)
	graph.Traverse() // traverse the graph
}
