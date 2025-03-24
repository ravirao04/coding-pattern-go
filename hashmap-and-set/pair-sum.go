package hashmap_and_set

func pairSum(arr []int, target int) []int {
	// k = a + b
	dic := map[int]int{}
	for index, item := range arr {
		dic[index] = item
	}
	for i, item := range arr {
		dif := target - item
		if index, ok := dic[dif]; ok {
			return []int{i, index}
		}
	}
	return []int{}
}
