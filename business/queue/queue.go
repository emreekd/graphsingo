package queue

import (
	"container/list"
	"../node"
)

type IQueue interface {
	Add(interface{})
	Pop()
	Size() int
}

type Queue struct {
	Items *list.List
}

func (q *Queue) Add(v *node.Node) {
	q.Items.PushBack(v)
}

func (q *Queue) Pop() *list.Element {
	var item = q.Items.Front()
	q.Items.Remove(item)
	return item
}
func (q *Queue) Size() int {
	return q.Items.Len()
}
