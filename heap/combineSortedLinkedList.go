package heap

import "container/heap"

type Node struct {
	Key  int
	Next *Node
}

type NodeItems []*Node

func (items NodeItems) Len() int {
	return len(items)
}
func (items NodeItems) Less(i, j int) bool {
	return items[i].Key > items[j].Key
}

func (items NodeItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items *NodeItems) Push(i any) {
	*items = append(*items, i.(*Node))
}
func (items *NodeItems) Pop() any {
	item := (*items)[len(*items)-1]
	*items = (*items)[0 : len(*items)-1]
	return item
}

type MinHeap struct {
	items *NodeItems
	size  int
}

func (min *MinHeap) Push(i any) {
	min.items.Push(i)
	min.size++
}

func (min *MinHeap) Pop() any {
	min.size--
	return min.items.Pop()
}

func (min *MinHeap) init() {
	heap.Init(min.items)
}

func (min *MinHeap) IsEmpty() bool {
	return min.size == 0
}

func (min *MinHeap) Fix() {
	heap.Fix(min.items, 0)
}

func combineSortedLinkedLists(linkedLists []*Node) *Node {
	dummy := &Node{Key: -1}
	root := dummy
	size := len(linkedLists)
	minHeap := &MinHeap{size: size}
	minHeap.init()

	for _, node := range linkedLists {
		minHeap.Push(node)
		//minHeap.Fix()
	}

	for !minHeap.IsEmpty() {
		node := minHeap.Pop().(*Node)
		dummy.Next = node
		dummy = dummy.Next
		if node.Next != nil {
			minHeap.Push(node.Next)
		}
	}

	return root.Next
}
