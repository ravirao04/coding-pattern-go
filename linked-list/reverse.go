package linked_list

func reverse(root *Node) *Node {
	if root.Next == nil {
		return root
	}
	newNode := reverse(root.Next)
	root.Next.Next = root
	root.Next = nil
	return newNode
}

func iterativeReverse(root *Node) *Node {
	var preNode *Node
	currentNode := root
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = preNode
		preNode = currentNode
		currentNode = nextNode
	}
	return preNode
}
