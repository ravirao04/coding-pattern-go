package stack

import "unicode"

func evaluateExpression(str string) int {
	res := 0
	sign := 1
	curr := 0
	stack := []int{}

	for _, char := range str {
		if unicode.IsDigit(char) {
			curr = curr*10 + int(char)
		} else if char == '+' || char == '-' {
			res = res + curr*sign
			curr = 0
			if char == '-' {
				sign = -1
			} else {
				sign = 1
			}
		} else if char == '(' {
			stack = append(stack, res, sign)
			res = 0
			sign = 1
		} else if char == ')' {
			res += sign * curr
			item1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res *= item1
			item2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += item2
			curr = 0
		}
	}
	return res + curr*sign
}
