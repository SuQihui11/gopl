package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "Febuaray", 3: "March", 12: "December"}

	fmt.Printf("Months: %v , type : %[1]T\n", months)

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Printf("Q2： %v, type : %[1]T\n", Q2) // ["April" "May" "June"]
	fmt.Println(summer)                      // ["June" "July" "August"]

	//fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	b := []int{1, 2, 3, 4, 5}
	c := []string{1: "a", 2: "b", 3: "c"}
	d := []string{1: "a", 2: "b", 3: "c"}

	fmt.Printf("type: b = %T, c = %T\n", b, c)
	str, _ := equal(c, d)
	fmt.Println(str)

	var s []int // len(s) == 0, s == nil
	s = nil     // len(s) == 0, s == nil
	// 所有的Go语言函数应该以相同的方式对待nil值的slice和0长度的slice!!!!
	// 如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。
	fmt.Println(len(s))
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	if s == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("not nil!")
	}

}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) (string, bool) {
	if len(x) != len(y) {
		return "", false
	}
	var str string
	for i := range x {
		str = fmt.Sprintf("type = %T, value = %[1]v\n", x[i])
		if x[i] != y[i] {
			return str, false
		}
	}
	return str, true
}
