package main

import (
	"fmt"
)

/**
	答案： BA

 */

func main() {
	//f := func() { fmt.Print("A") }
	//defer f()
	//f = func() { fmt.Print("B") }
	//defer f()
	//
	//f1 := func(a int) {
	//	 fmt.Println("a")
	//}
	//defer f1(getNumber())


	//rand.Seed(time.Now().Unix())
	//var tmpSlic []tmp
	//for i:=0; i < 100; i++ {
	//	tmp := tmp{
	//		Name:fmt.Sprintf("%d", i),
	//		Status: rand.Int31n(4)+2,
	//	}
	//	tmpSlic = append(tmpSlic, tmp)
	//}
	//
	//sort.Slice(tmpSlic, func(i, j int) bool {
	//	if tmpSlic[i].Status == 2 {
	//		return true
	//	}
	//	if tmpSlic[i].Status == 3 {
	//		return false
	//	}
	//	if tmpSlic[j].Status == 3 {
	//		return true
	//	}
	//	return tmpSlic[i].Status < tmpSlic[j].Status
	//})
	//for _, val := range tmpSlic {
	//	fmt.Printf("%+v\n", val)
	//}

	v := [...]int{1:2, 3:5}
	fmt.Println(len(v))

	var i float64 = 3/2
	fmt.Print(i)
}

func getNumber() int {

	fmt.Println("C")
	return 2
}

type tmp struct {
	Name string
	Status int32
}


