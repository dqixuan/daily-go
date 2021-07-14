package main

import "fmt"

func main() {
	aa := a{}
	var bb t
	bb = &b{aa}
	if v, ok := bb.(t); ok {
		fmt.Println(v.Test())
	}
	fmt.Println(bb.Test())
}

type a struct {
}

type b struct {
	a
}

type t interface {
	Test() string
}

func (a *a) Test() string {

	return "test"
}
