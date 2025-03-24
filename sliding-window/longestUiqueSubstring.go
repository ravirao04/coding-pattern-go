package sliding_window

import "math"

func longestUniqueSubString(str string) int {
	maxLength := 0
	freq := make(map[rune]struct{})
	left, right := 0, 0
	for right < len(str) {
		char := rune(str[right])
		if _, found := freq[char]; found {
			delete(freq, char)
			left++
		}
		maxLength = int(math.Max(float64(maxLength), float64(right-left+1)))
		freq[char] = struct{}{}
		right++
	}
	return maxLength
}

func optimisedLongestUniqueSubString(str string) int {
	maxLength := 0
	freq := make(map[rune]int)
	left, right := 0, 0
	for right < len(str) {
		char := rune(str[right])
		if index, found := freq[char]; found {
			delete(freq, char)
			left = index + 1
		}
		maxLength = int(math.Max(float64(maxLength), float64(right-left+1)))
		freq[char] = right
		right++
	}
	return maxLength
}
