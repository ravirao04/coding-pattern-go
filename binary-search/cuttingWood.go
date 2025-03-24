package binary_search

import "math"

// calculate height to be cut
func cuttingWood(arr []int, k int) int {
	left, right := 0, int(math.Max(arr))
	for left < right {
		mid := right - left/2 + 1
		if FindHeightWoodCollected(arr, mid, k) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
func FindHeightWoodCollected(arr []int, height, target int) bool {
	sum := 0
	for _, item := range arr {
		if item > height {
			sum = (item - target) + sum
		}
	}
	return sum >= target
}
