package sliding_window

import "math"

func LongestUniformSubstring(str string, k int) int {
	maxFreq := 0
	maxLength := 0
	left, right := 0, 0
	frq := make(map[rune]int)
	for right < len(str) {
		char := rune(str[right])
		frq[char] += 1
		maxFreq = int(math.Max(float64(maxFreq), float64(frq[char])))
		charToChange := (right - left + 1) - maxFreq
		if charToChange > k {
			frq[rune(str[left])] -= 1
			left++
		}
		maxLength = right - left + 1
		right++
	}
	return maxLength
}
