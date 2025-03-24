package binary_search

func findTargetInRotatedSortedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := right - left/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			if arr[left] < target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if arr[mid] < target && target < arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if arr[left] == target {
		return left
	}
	return -1
}
