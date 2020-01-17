package node

type INode interface {
	AddAdjacent(Node)
}
type Node struct {
	Value     int
	Adjacents []Node
}

func (n *Node) AddAdjacent(node Node) {
	n.Adjacents = append(node.Adjacents, node)
}
