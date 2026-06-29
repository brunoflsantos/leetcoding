package validparentheses

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := 0
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack[top] = r
			top++
		case ')':
			if top > 0 && stack[top-1] == '(' {
				top--
			} else {
				return false
			}
		case ']':
			if top > 0 && stack[top-1] == '[' {
				top--
			} else {
				return false
			}
		case '}':
			if top > 0 && stack[top-1] == '{' {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}
