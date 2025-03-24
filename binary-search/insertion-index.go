package binary_search

import "sort"

func insertionIndex(arr []int, target int) int {
	sort.Ints(arr)
	left, right := 0, len(arr)
	for left < right {
		mid := (right - left) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}
