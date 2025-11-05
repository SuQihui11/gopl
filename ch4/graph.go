package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	class := make(map[string]int, 4)
	input := bufio.NewReader(os.Stdin)
	for {
		r, _, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}

		t := whichType(r)
		class[t]++
		counts[r]++
	}
	fmt.Printf("%#v", class)
}

func whichType(c rune) string {
	if unicode.IsLetter(c) {
		return "letter"
	} else if unicode.IsDigit(c) {
		return "digit"
	} else if unicode.Is(unicode.Han, c) {
		return "han"
	} else {
		return "unknown"
	}
}

//// 判断一个unicode码点为什么类型
//func whichType(r rune) string {
//	if unicode.IsNumber(r) {
//		return "number"
//	}
//	if unicode.Is(unicode.Han, r) {
//		return "中文"
//	}
//	if r >= 65 && r <= 122 {
//		return "English"
//	}
//
//	return "any"
//}
