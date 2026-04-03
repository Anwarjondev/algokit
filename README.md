# algokit

A small Go toolkit with generic, test-covered implementations of:

- `stack`: LIFO stack
- `queue`: FIFO queue

## Requirements

- Go `1.24.4` (or compatible)

## Project structure

```text
algokit/
├── go.mod
├── stack/
│   ├── stack.go
│   └── stack_test.go
└── queue/
    ├── queue.go
    └── queue_test.go
```

## Run tests

From the project root:

```bash
go test ./...
```

## Stack

Package: `algokit/stack`

### API

- `New[T any]() *Stack[T]`
- `(*Stack[T]).Push(value T)`
- `(*Stack[T]).Pop() (T, bool)`
- `(*Stack[T]).Peek() (T, bool)`
- `(*Stack[T]).Len() int`
- `(*Stack[T]).IsEmpty() bool`

### Example

```go
package main

import (
	"fmt"

	"algokit/stack"
)

func main() {
	s := stack.New[int]()
	s.Push(10)
	s.Push(20)

	if top, ok := s.Peek(); ok {
		fmt.Println("peek:", top) // 20
	}

	if v, ok := s.Pop(); ok {
		fmt.Println("pop:", v) // 20
	}

	fmt.Println("len:", s.Len())         // 1
	fmt.Println("isEmpty:", s.IsEmpty()) // false
}
```

## Queue

Package: `algokit/queue`

### API

- `New[T any]() *Queue[T]`
- `(*Queue[T]).Enqueue(val T)`
- `(*Queue[T]).Dequeue() (T, bool)`
- `(*Queue[T]).Front() (T, bool)`
- `(*Queue[T]).Back() (T, bool)`
- `(*Queue[T]).Size() int`
- `(*Queue[T]).IsEmpty() bool`

### Example

```go
package main

import (
	"fmt"

	"algokit/queue"
)

func main() {
	q := queue.New[string]()
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	if f, ok := q.Front(); ok {
		fmt.Println("front:", f) // a
	}

	if v, ok := q.Dequeue(); ok {
		fmt.Println("dequeue:", v) // a
	}

	if b, ok := q.Back(); ok {
		fmt.Println("back:", b) // c
	}

	fmt.Println("size:", q.Size())       // 2
	fmt.Println("isEmpty:", q.IsEmpty()) // false
}
```
