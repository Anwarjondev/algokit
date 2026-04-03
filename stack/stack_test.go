package stack

import "testing"

func TestStack_PushPop(t *testing.T) {
	s := New[int]()

	s.Push(10)
	s.Push(20)

	val, ok := s.Pop()
	if !ok || val != 20 {
		t.Errorf("expected 20, got %v", val)
	}

	val, ok = s.Pop()
	if !ok || val != 10 {
		t.Errorf("expected 10, got %v", val)
	}

	_, ok = s.Pop()
	if ok {
		t.Errorf("expected empty stack")
	}
}

func TestStack_Peek(t *testing.T) {
	s := New[int]()

	s.Push(5)
	s.Push(15)

	val, ok := s.Peek()
	if !ok || val != 15 {
		t.Errorf("expected 15, got %v", val)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Errorf("expected empty stack")
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Errorf("expected non-empty stack")
	}
}
func TestStack_String(t *testing.T) {
	s := New[string]()

	s.Push("hello")
	s.Push("world")

	val, _ := s.Pop()
	if val != "world" {
		t.Errorf("expected world, got %v", val)
	}
}