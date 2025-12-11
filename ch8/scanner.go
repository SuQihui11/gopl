package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. 读取一个文件
	file, err := os.OpenFile("./test.txt", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 2.读取一个用户的输入
	scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 3.读取一个字符串的输入
	str := "hello world"
	reader := strings.NewReader(str)
	scanner = bufio.NewScanner(reader)

	// 按照单词分割
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
