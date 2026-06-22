type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	var zero T
	if s.IsEmpty() {
		return zero
	}
	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx]
	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}


func isValid(s string) bool {
	stk := NewStack[byte]()
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stk.Push(c)
			continue
		}

		if stk.IsEmpty() {
			return false
		}
		top := stk.Pop()

		if (c == ')' && top != '(') ||
			(c == '}' && top != '{') ||
			(c == ']' && top != '[') {
			return false
		}
	}

	return stk.IsEmpty()
}