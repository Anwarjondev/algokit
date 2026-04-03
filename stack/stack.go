package stack


type Stack[T any] struct {
	stack []T
}


func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Len() == 0 {
		var zeroVal T
		return zeroVal, false
	}
	val := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Len() == 0 {
		var zeroVal T
		return zeroVal, false
	}
	return s.stack[s.Len()-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}