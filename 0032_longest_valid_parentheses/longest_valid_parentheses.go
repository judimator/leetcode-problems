package longest_valid_parentheses

type stack struct {
	chars []rune
}

func (s *stack) push(char rune) {
	s.chars = append(s.chars, char)
}

func (s *stack) pop() rune {
	item := s.chars[len(s.chars)-1]
	s.chars = s.chars[0 : len(s.chars)-1]

	return item
}

func (s *stack) len() int {
	return len(s.chars)
}

func isValid(s string) bool {
	stack := &stack{}

	for i, char := range s {
		if i == 0 && char == ')' {
			return false
		}

		if char == '(' {
			stack.push(char)
		} else {
			if stack.len() == 0 {
				return false
			}

			stack.pop()
		}
	}

	return stack.len() == 0
}

func longestValidParentheses2(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isValid(s[i : j+1]) {
				if (j - i) > result {
					result = j - i + 1
				}
			}
		}
	}

	return result
}

func longestValidParentheses(s string) int {
	stack := []int{-1}
	result := 0

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = maximum(result, i-stack[len(stack)-1])
			}
		}
	}

	return result
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
