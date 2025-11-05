package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
func basename1(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func common(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return common(s[:n-3]) + "," + s[n-3:] //递归
}

func common1(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		// 两种情况：1.均分 2.有剩余 ，但是无论如何都是需要在最开始就写入的
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func common2(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	if s[0] == '+' || s[0] == '-' {
		start = 1
	}
	idx := strings.Index(s, ".")
	intStr := common1(s[start:idx])
	if start == 0 {
		return intStr + s[idx:]
	} else {
		return string(s[0]) + intStr + s[idx:]
	}
}

func isSame(a, b string) bool {
	l1, l2 := len(a), len(b)
	if l1 != l2 || a == b {
		return false
	}

	m1 := make(map[rune]int)

	for _, v := range a {
		m1[v]++
	}
	for _, v := range b {
		m1[v]--
	}

	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s2)
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "a"))
	fmt.Println(strings.Count(s, "a"))
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.HasSuffix(s, "a"))
	fmt.Println(strings.Index(s, "a"))
	fmt.Println(strings.HasPrefix(s, "a"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))

	fmt.Println(common2("1234567.2222"))
	fmt.Println(common2("-1234567.2222"))

	//x := 123
	//y := fmt.Sprintf("%d", x)
	//fmt.Println(y, strconv.Itoa(x))             // "123 123"
	//fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	//s = fmt.Sprintf("x=%b", x)                  // "x=1111011"
	//fmt.Println(s)
	//x, err := strconv.Atoi("123")             // x is an int
	//y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

}
