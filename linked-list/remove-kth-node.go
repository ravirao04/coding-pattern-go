package linked_list

func removeKthNode(root *Node, kth int) *Node {
	DummyNode := &Node{key: -1, Next: root}
	leader, head := DummyNode, DummyNode
	for i := 0; i < kth; i++ {
		if leader == nil {
			return root
		}
		leader = leader.Next
	}

	for leader.Next != nil {
		head = head.Next
		leader = leader.Next
	}

	head.Next = head.Next.Next
	return DummyNode.Next

}
