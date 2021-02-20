package stack

type IStackIterator interface {
	More() bool
	Next() (error, interface{})
}
