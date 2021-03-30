package main

import "fmt"

// 实现观察者模式， 使用场景Redis pub/sub

type observer interface {
	Update(event string)
}

type publisher struct {
	observers []observer
}

func (pu *publisher) AddObserver(o observer) {
	pu.observers = append(pu.observers, o)
}

func (pu *publisher) Publish(event string) {
	for _, obj := range pu.observers {
		obj.Update(event)
	}
}

type AObserver struct {
}

func (a *AObserver) Update(event string) {
	fmt.Println("A recieved event ", event)
}

type BObserver struct {
}

func (a *BObserver) Update(event string) {
	fmt.Println("B recieved event ", event)
}

func main() {
	pb := &publisher{}
	a := &AObserver{}
	b := &BObserver{}
	pb.AddObserver(a)
	pb.AddObserver(b)

	pb.Publish("hello")
}
