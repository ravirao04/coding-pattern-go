package fast_and_slow_pointer

import "math"

func isHappyNumber(k int) bool {
	slow, fast := k, k
	for {
		slow := getNext(slow)
		fast := getNext(getNext(fast))
		if fast == 1 {
			return true
		} else if slow == fast {
			return false
		}
	}
	return false
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		rem := n % 10
		n = n / 10
		sum += int(math.Pow(float64(rem), 2.0))
	}
	return sum
}
