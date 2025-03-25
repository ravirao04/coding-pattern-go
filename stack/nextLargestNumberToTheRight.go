package stack

func nextLargestNumberToTheRight(arr []int) []int {
	n := len(arr)
	stack := make([]int, n)
	result := make([]int, n)

	for i := len(arr) - 1; i <= 0; i++ {
		item := arr[i]

		for len(stack) != 0 && stack[len(stack)-1] < item {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = -1
		} else {
			peek := stack[len(stack)-1]
			result[i] = peek
		}
		stack = append(stack, item)
	}
	return result
}
