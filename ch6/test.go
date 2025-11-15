package main

//type p *int
//
//func (p) f() {
//
//}

type IntList struct {
	value int
	tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.tail.Sum()
}
