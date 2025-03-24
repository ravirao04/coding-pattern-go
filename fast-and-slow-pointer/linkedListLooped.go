package fast_and_slow_pointer

func isLinkedListLooped(root *Node) bool {
	fast, slow := root, root
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
