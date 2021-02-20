package stack

import "errors"

type MyStack struct {
	items    []interface{}
	capacity int
	top      int
	version  int64
}

var emptyStackError = errors.New("empty stack")

func NewArrayStack(capacity int) IStack {
	if capacity < 0 {
		capacity = 0
	}
	return &MyStack{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		top:      -1,
	}
}

func (ms *MyStack) Size() int {
	return ms.top + 1
}

func (ms *MyStack) IsEmpty() bool {
	return ms.top <= 0
}

func (ms *MyStack) Push(value interface{}) {
	if ms.top+1 == ms.capacity {
		newItems := make([]interface{}, ms.capacity*2)
		copy(newItems, ms.items)
		ms.items = newItems
		ms.capacity *= 2
	}
	ms.top++
	ms.items[ms.top] = value
}
