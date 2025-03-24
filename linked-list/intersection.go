package linked_list

func isIntersected(aNode, bNode *Node) *Node {
	ptrA, ptrB := aNode, bNode
	for ptrA != ptrB {
		ptrA = ptrA.Next
		ptrB = ptrB.Next
		if ptrA == nil {
			ptrA = bNode
		}
		if ptrB == nil {
			ptrB = aNode
		}
	}
	return ptrA
}
