package main

import "fmt"

/**

 */

type Call interface {
	Run()
}

type Cat struct {
}

func (c Cat) Run() {
	fmt.Println("cat run")
}

type Dog struct {
}

func(d Dog) Run() {
	fmt.Println("dog run")
}

type Animal struct {
	call Call
	Name string
}

func main() {
	a := Animal{
		call: Dog{},
		Name: "a" ,
	}
	a.call.Run()
	a.call = Cat{}
	a.call.Run()
}
