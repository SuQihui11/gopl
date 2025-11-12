package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.kimi.com"
	err := outlinePractice(url)
	if err != nil {
		log.Fatal(err)
	}
}

func outlinePractice(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	// 匿名函数也就是闭包的意义之一就在于，访问外部的局部变量，减少函数的入参数量
	var deep int

	start := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", deep*2, "", n.Data)
			deep++
		}
	}

	end := func(n *html.Node) {
		if n.Type == html.ElementNode {
			deep--
			fmt.Printf("%*s</%s>\n", deep*2, "", n.Data)
		}
	}

	forEach(doc, start, end)
	return nil
}
