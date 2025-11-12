package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 模拟 resp.Request.URL
	baseURL, _ := url.Parse("https://golang.org/doc/code.html")

	// 测试不同的href值
	testURLs := []string{
		"/pkg/",
		"install.html",
		"../project/",
		"https://external.com",
	}

	for item := range testURLs {
		link, _ := baseURL.Parse(testURLs[item])
		fmt.Println(link)
	}

}
