package validparentheses

func isValid(s string) bool {
	parenthesesMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, len(s))
	top := -1
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			top++
			stack[top] = r
		} else if p := parenthesesMap[r]; top >= 0 && stack[top] == p {
			top--
		} else {
			return false
		}
	}
	return top == -1
}
