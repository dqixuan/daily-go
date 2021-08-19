package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	dv := &doubleValue{}
	i := 0
	for key, val := range m {
		if i < k {
			dv.Push([2]int{key, val})
		} else {
			heap.Init(dv)
			if (*dv)[0][1] < val {
				(*dv)[0] = [2]int{key, val}
				heap.Init(dv)
			}
		}
		i++
	}
	var ans []int
	for i:=0; i < dv.Len(); i++ {
		ans = append(ans, (*dv)[i][0])
	}
	return ans
}



type doubleValue [][2]int

func(dv doubleValue) Len() int {
	return len(dv)
}

func (dv doubleValue) Less(i, j int) bool {
	return dv[i][1] < dv[j][1]
}

func (dv doubleValue) Swap(i, j int) {
	dv[i], dv[j] = dv[j], dv[i]
}

func (dv *doubleValue) Push(i interface{}) {
	*dv = append(*dv, i.([2]int))
}

func (dv *doubleValue) Pop() interface{} {
	if dv.Len() == 0 {
		return nil
	}
	e := (*dv)[dv.Len()-1]
	*dv = (*dv)[:dv.Len()-1]
	return e
}


