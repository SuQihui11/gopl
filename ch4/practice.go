package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	//s := []string{"a", "a", "a", "b", "c", "d"}
	//s = removeDuplicates(s)
	//fmt.Printf("%q\n", s)
	//
	//str := "hello    work"
	//runes := []byte(str)
	//r, size := utf8.DecodeRune(runes[6:])
	//fmt.Printf("%q\n", r)
	//fmt.Printf("%v\n", size)
	//
	//// 测试用例
	//tests := []string{
	//	"hello  world",
	//	"a   b   c",
	//	"  leading spaces",
	//	"trailing spaces  ",
	//	"  multiple   spaces   everywhere  ",
	//	"hello\t\tworld",      // 制表符
	//	"hello\n\nworld",      // 换行符
	//	"你好  世界",              // 中文字符
	//	"mixed  \t\n  spaces", // 混合空白字符
	//	"noSpaces",
	//	"",
	//}
	//
	//for _, test := range tests {
	//	input := []byte(test)
	//	result := CompressSpaces(input)
	//	fmt.Printf("输入: %q\n输出: %q\n\n", test, string(result))
	//}
	//
	//arr := []int{1, 2, 3, 4, 5, 6}
	//RotateLeftReverse(arr, 2)
	//fmt.Printf("%v\n", arr)

	// 示例 1: 包含 ASCII 和多字节字符
	s1 := []byte("Hello, 世界")
	fmt.Printf("原始 (s1): %s\n", s1)
	fmt.Printf("原始 (hex): %x\n", s1)

	s1 = reverseBytes(s1)

	fmt.Printf("反转 (s1): %s\n", s1)
	fmt.Printf("反转 (hex): %x\n", s1)
	// 预期输出: "界世 ,olleH"

	fmt.Println(string(make([]byte, 20))) // 分隔符

	// 示例 2: 只有多字节字符
	s2 := []byte("Go语言")
	fmt.Printf("原始 (s2): %s\n", s2)
	fmt.Printf("原始 (hex): %x\n", s2)

	s2 = reverseBytes(s2)

	fmt.Printf("反转 (s2): %s\n", s2)
	fmt.Printf("反转 (hex): %x\n", s2)
	// 预期输出: "言语oG"
}

// 反转数组，使用的是数组指针
const length = 10

func reverseArray(s *[length]int) {
	if s == nil {
		return
	}
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//编写一个rotate函数，通过一次循环完成旋转

// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
func removeDuplicates(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	for i := 1; i < len(s); {
		if s[i] == s[i-1] {
			// 复制到这里进行处理
			copy(s[i:], s[i+1:])
			// 删除最后一位
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func CompressSpaces(str []byte) []byte {
	if len(str) == 0 {
		return str
	}

	writePos := 0
	readPos := 0

	lastSpace := false
	for readPos < len(str) {
		r, size := utf8.DecodeRune(str[readPos:])
		if unicode.IsSpace(r) {
			// 是空格
			if !lastSpace {
				str[writePos] = ' '
				writePos++
				lastSpace = true
			}
			readPos += size
		} else {
			copy(str[writePos:writePos+size], str[readPos:readPos+size])
			writePos += size
			readPos += size
			lastSpace = false
		}
	}
	return str[:writePos]
}

// RotateLeftReverse 使用三次反转法左旋转
func RotateLeftReverse(s []int, k int) {
	n := len(s)
	k = k % n
	if n == 0 || k == 0 {
		return
	}
	// 首先全部旋转
	reverse2(s, 0, n-1)
	// 旋转后k部分
	reverse2(s, n-k, n-1)
	reverse2(s, 0, n-k-1)
	return
}

// reverse 反转数组的指定区间
func reverse2(s []int, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
func ReverseUTF8(str []byte) []byte {
	runes := []rune(string(str))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
func reverseBytes(slice []byte) []byte {
	var idx = 0
	for idx < len(slice) {
		_, size := utf8.DecodeRune(slice[idx:])
		for i, j := idx, idx+size-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		idx += size
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
