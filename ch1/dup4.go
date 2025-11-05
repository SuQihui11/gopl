package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNameMap := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "Please give me a file name!")
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				continue
			}
			countLines(file, counts, fileNameMap)
			_ = file.Close()
		}
	}
	for name, count := range counts {
		if count > 1 {
			fmt.Printf("重复行：%s，出现重复行的文件名称：%s，重复次数：%d\r\n", name, fileNameMap[name], count)
		}
	}
}

func countLines(file *os.File, counts map[string]int, nameMap map[string]string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
		nameMap[input.Text()] = file.Name()
	}
}
