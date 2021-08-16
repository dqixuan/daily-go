package stack

type IStack interface {
	// 元素个数
	Size() int
	// 栈是否为空
	IsEmpty() bool

	// 添加元素
	Push(value interface{})
	// 取栈顶元素
	Pop() (interface{}, error)
	// 查看栈顶元素
	Peek() (interface{}, error)

	String() string
}
