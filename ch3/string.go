package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	fmt.Printf("%T %[1]s\n", s)
	fmt.Printf("%T %[1]v\n", s[0])

	//c := s[len(s)] // panic: index out of range
	//s[0] = 'L' // compile error: cannot assign to s[0]
	s1 := "世界"
	s2 := "\xe4\xb8\x96\xe7\x95\x8c"
	s3 := "\u4e16\u754c"
	s4 := "\U00004e16\U0000754c"

	fmt.Println(s1 == s2) // true
	fmt.Println(s2 == s3) // true
	fmt.Println(s3 == s4) // true

	fmt.Println(s1) // 世界
	fmt.Println(s2) // 世界
	fmt.Println(s3) // 世界
	fmt.Println(s4) // 世界

	s = "世"

	// 1. Unicode 码点
	r := '界'
	fmt.Printf("Unicode 码点: U+%04X (%d)\n", r, r)
	// 输出: Unicode 码点: U+4E16 (20022)

	// 2. UTF-8 编码（字节）
	fmt.Printf("UTF-8 字节: %x\n", []byte(s))
	// 输出: UTF-8 字节: e4 b8 96

	// 3. 字节数
	fmt.Printf("字节数: %d\n", len(s))
	// 输出: 字节数: 3

	// 4. 字符数
	fmt.Printf("字符数: %d\n", utf8.RuneCountInString(s))
	// 输出: 字符数: 1

	s = "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// "program" in Japanese katakana
	s = "プログラム"
	fmt.Printf("%d bytes, %d characters \n", len(s), utf8.RuneCountInString(s))
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r1 := []rune(s)
	fmt.Printf("%x\n", r1) // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Printf("%s\n", string(r1))

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

	fmt.Println(string(1234567)) // "?"

}
