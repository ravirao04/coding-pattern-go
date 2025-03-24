package tow_pointer

import "math"

func LargeContainer(arr []int) int {
	left, right := 0, len(arr)-1
	maxWater := 0
	for left < right {
		water := calculateWater(left, right, arr[left], arr[right])
		maxWater = int(math.Max(float64(water), float64(maxWater)))
		if arr[left] < arr[right] {
			left++
		} else if arr[left] > arr[right] {
			right--
		} else {
			left++
			right++
		}
	}
	return maxWater
}

func calculateWater(aWidth, bWidth, aHeight, bHeight int) int {
	return int(math.Min(float64(aHeight), float64(bHeight)) * float64(bWidth-aWidth))
}
