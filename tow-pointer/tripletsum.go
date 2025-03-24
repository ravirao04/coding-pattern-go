package tow_pointer

func TripletSum(arr []int) [][]int {
	// a+b+c = 0
	result := [][]int{}
	for i := 0; i < len(arr); i++ {
		// a = - (b+c)
		pairs := pairsum(arr[i:], -arr[i])
		for _, items := range pairs {
			result = append(result, append(items, i))
		}
	}
	return result
}

func pairsum(arr []int, target int) [][]int {
	result := [][]int{}
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			result = append(result, []int{left, right})
			left++
		}
		for left < right && arr[left] == arr[left-1] {
			left++
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return result
}
