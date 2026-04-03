package queue

import "testing"

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, ok := q.Dequeue()
	if !ok || v != 1 {
		t.Errorf("expected 1, got %v", v)
	}

	v, ok = q.Dequeue()
	if !ok || v != 2 {
		t.Errorf("expected 2, got %v", v)
	}

	v, ok = q.Dequeue()
	if !ok || v != 3 {
		t.Errorf("expected 3, got %v", v)
	}

	_, ok = q.Dequeue()
	if ok {
		t.Errorf("expected empty queue")
	}
}


func TestQueue_FrontBack(t *testing.T) {
	q := New[int]()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	front, _ := q.Front()
	if front != 10 {
		t.Errorf("expected front 10, got %v", front)
	}

	back, _ := q.Back()
	if back != 30 {
		t.Errorf("expected back 30, got %v", back)
	}
}

func TestQueue_SizeIsEmpty(t *testing.T) {
	q := New[int]()

	if !q.IsEmpty() {
		t.Errorf("expected empty queue")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	q.Dequeue()
	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("expected empty queue after dequeues")
	}
}


func TestQueue_String(t *testing.T) {
	q := New[string]()

	q.Enqueue("a")
	q.Enqueue("b")

	v, _ := q.Dequeue()
	if v != "a" {
		t.Errorf("expected 'a', got %v", v)
	}
}

func TestQueue_Stress(t *testing.T) {
	q := New[int]()

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 1000; i++ {
		v, ok := q.Dequeue()
		if !ok || v != i {
			t.Errorf("expected %d, got %v", i, v)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue after stress test")
	}
}


func TestQueue_MixedOperations(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)

	v, _ := q.Dequeue()
	if v != 1 {
		t.Errorf("expected 1, got %v", v)
	}

	q.Enqueue(3)

	v, _ = q.Dequeue()
	if v != 2 {
		t.Errorf("expected 2, got %v", v)
	}

	v, _ = q.Dequeue()
	if v != 3 {
		t.Errorf("expected 3, got %v", v)
	}
}