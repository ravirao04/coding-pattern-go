package stack

func validateParanthesisExpression(str string) bool {
	stack := make([]rune, len(str))

	openDic := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range str {
		if _, found := openDic[char]; found {
			stack = append(stack, char)
		} else {
			if len(stack) > 0 && openDic[stack[len(stack)-1]] == char {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
