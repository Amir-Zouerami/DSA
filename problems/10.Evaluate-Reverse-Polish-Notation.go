package leetcode

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))

	operationMappings := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	result, _ := strconv.Atoi(tokens[len(tokens)-1])

	for i := range tokens {
		if operation, ok := operationMappings[tokens[i]]; ok {
			popped1, popped2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			intPopped1, _ := strconv.Atoi(popped1)
			intPopped2, _ := strconv.Atoi(popped2)

			result = operation(intPopped1, intPopped2)

			if i+1 == len(tokens) {
				break
			}

			stack = append(stack, strconv.Itoa(result))
		} else {
			stack = append(stack, tokens[i])
		}
	}

	return result
}
