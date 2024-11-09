package valid_parentheses

type stack struct {
	chars []rune
}

func isValid(s string) bool {
	stack := &stack{}

	for _, char := range s {
		if char == '[' || char == '(' || char == '{' {
			stack.push(char)
		} else {
			if stack.len() == 0 {
				return false
			}

			c := stack.pop()

			if c == '[' {
				if char != ']' {
					return false
				}
			} else if c == '(' {
				if char != ')' {
					return false
				}
			} else {
				if char != '}' {
					return false
				}
			}
		}
	}

	return stack.len() == 0
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
