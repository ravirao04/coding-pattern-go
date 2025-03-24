package binary_search

func FirstAndLastOccurance(arr []int, k int) []int {
	lower := lowerBound(arr, k)
	upper := upperBound(arr, k)
	return []int{lower, upper}
}

func upperBound(arr []int, k int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := right - left/2
		if arr[mid] == k {
			return mid
		} else if arr[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if arr[left] == k {
		return left
	}
	return -1
}

func lowerBound(arr []int, k int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := right - left/2 + 1 // right bias
		if arr[mid] == k {
			return mid
		} else if arr[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if arr[left] == k {
		return left
	}
	return -1
}
