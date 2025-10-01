package top100

func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if (s[i] == '}' && stack[len(stack)-1] != '{') || (s[i] == ']' && stack[len(stack)-1] != '[') || (s[i] == ')' && stack[len(stack)-1] != '(') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
