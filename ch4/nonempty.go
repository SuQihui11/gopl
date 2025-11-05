// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2)) // `["one" "three"]`
	fmt.Printf("%q\n", data2)

	s := []int{5, 6, 7, 8, 9}
	//fmt.Println(remove(s, 2)) // "[5 6 8 9]"
	s = remove2(s, 2)
	fmt.Println(s)
	fmt.Printf("%v\n", len(s))
	fmt.Printf("%v\n", s)

	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	top := stack[len(stack)-1]
	// 收缩
	stack = stack[:len(stack)-1]
	fmt.Printf("%v\n", top)
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
