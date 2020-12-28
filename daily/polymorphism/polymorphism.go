package main

import "fmt"

/*
多态：golang通过接口实现多态
某个类实现接口定义的所有方法，那么它就实现了该接口。
利用该接口实现多态
*/

// Income 收入接口
type Income interface {
	calculate()  int 
	source() string
}

// FixedBilling 项目名称
type FixedBilling struct {
	projectName string  // 项目名称
	biddedAmount int    // 投标该项目的金额
}

// TimeAndMaterial name
type TimeAndMaterial struct {
	projectName string
	noOfHours int
	hourRate  int
}

func (f *FixedBilling) calculate() int {
	return f.biddedAmount
}

func (f *FixedBilling) source() string {
	return f.projectName
}

func (tm *TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourRate
}

func (tm *TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateIncome(ic []Income) {
	var netIncome int
	for _, v := range ic {
		fmt.Printf("Income From %s = $%d\n", v.source(), v.calculate())
		netIncome += v.calculate()
	}
	fmt.Printf("Net income of organisation = $%d\n", netIncome)
}


func main() {
	f := &FixedBilling{
		projectName: "projectF",
		biddedAmount: 200,
	}
	tm := &TimeAndMaterial{
		projectName: "projectTM",
		noOfHours: 20,
		hourRate:  20,
	}
	incomeSlice := []Income{}
	incomeSlice = append(incomeSlice, f)
	incomeSlice = append(incomeSlice, tm)
	calculateIncome(incomeSlice)
}