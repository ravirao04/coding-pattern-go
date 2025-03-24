package tow_pointer

// time O(n) space O(1)
func PairSum(arr []int, target int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right++
		}
	}
	return []int{}
}

func BruteForcePairSum(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
