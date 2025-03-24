package fast_and_slow_pointer

func midPoint(root *Node) *Node {
	slow, fast := root, root
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
