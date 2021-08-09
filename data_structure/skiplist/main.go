package main

import "fmt"

func main() {
	sl := NewSkipList(5)
	sl.Add(2, 10)
	sl.Add(3, 20)

	v, ok := sl.Search(2)
	fmt.Println("v=", v, ", ok=", ok)

	v, ok = sl.Search(3)
	fmt.Println("v=", v, ", ok=", ok)

	sl.Remove(3)

	v, ok = sl.Search(3)
	fmt.Println("v=", v, ", ok=", ok)

}

