package leetcode

func IsValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := make([]rune, 0, len(s))
	relations := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	for _, letter := range s {
		if letter == '(' || letter == '{' || letter == '[' {
			stack = append(stack, letter)
		} else {
			if len(stack) == 0 {
				return false
			}

			if relations[stack[len(stack)-1]] != letter {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
