package graph

import (
	Node "../node"
	queue "../queue"
	"container/list"
	"fmt"
)

type IGraph interface {
	Traverse()
	Add(n int, w int)
	GetNodes() *[]Node.Node
}

type BFSGraph struct {
	Nodes *[]Node.Node
}
type DFSGraph struct {
	Nodes *[]Node.Node
}

func(g *BFSGraph) GetNodes() *[]Node.Node{
	return g.Nodes;
}
func(g *DFSGraph) GetNodes() *[]Node.Node{
	return g.Nodes;
}

func (g *BFSGraph) Add(n int, w int) {
	addToGraph(n,w,g)
}

func (g *DFSGraph) Add(n int, w int) {
	addToGraph(n,w,g)
}
func addToGraph(n int, w int, g IGraph){
	que := &queue.Queue{
		Items: list.New(),
	}
	que.Add(&(*g.GetNodes())[0])
	for que.Size() > 0 {
		var queuedNode = que.Pop()
		var node = (*(&queuedNode.Value)).(*Node.Node)
		if node.Value == n {
			newNode := &Node.Node{
				Value: w,
			}
			node.Adjacents = append(node.Adjacents, *newNode)
		}else {
			if node.Adjacents != nil && len(node.Adjacents) > 0{
				for i := 0; i < len(node.Adjacents); i++{
					que.Add(&(node.Adjacents[i]))
				}
			}
		}
	}
}

func (g *BFSGraph) Traverse() {
	var visited map[int]bool
	visited = make(map[int]bool)
	que := &queue.Queue{
		Items: list.New(),
	}
	que.Add(&(*g.Nodes)[0])
	for que.Size() > 0 {
		var queuedNode = que.Pop()
		var node = (*(&queuedNode.Value)).(*Node.Node)
		if  !visited[node.Value]{
			fmt.Println(node.Value)
			visited[node.Value] = true
		}
		for i := 0; i < len(node.Adjacents); i++{
			que.Add(&(node.Adjacents[i]))
		}
	}
}


func (g *DFSGraph) Traverse() {
	var visited map[int]bool
	visited = make(map[int]bool)
	que := &queue.Queue{
		Items: list.New(),
	}
	que.Add(&(*g.Nodes)[0])
	for que.Size() > 0 {
		var queuedNode = que.PopFromBack()
		var node = (*(&queuedNode.Value)).(*Node.Node)
		if  !visited[node.Value]{
			fmt.Println(node.Value)
			visited[node.Value] = true
		}
		for i := 0; i < len(node.Adjacents); i++{
			que.Add(&(node.Adjacents[i]))
		}
	}
}

