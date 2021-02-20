package stack

type IStack interface {
	Size() int
	IsEmpty() bool

	Push(value interface{})
	Pop() (error, interface{})
	Peek() (error, interface{})

	Iterator() IStackIterator
	String() string
}
