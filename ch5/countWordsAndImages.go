package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages("https://www.kimi.com/kimiplus/cu52bqh7l5gqdkncdg01")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}

func CountWordsAndImages(url string) (int, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	words, images := countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		strArr := strings.Split(n.Data, " ")
		words += len(strArr)
	} else if n.Data == "img" {
		images++
	}
	var wc, ic int
	if n.Data != "style" || n.Data != "script" {
		wc, ic = countWordsAndImages(n.FirstChild)
	}
	wb, ib := countWordsAndImages(n.NextSibling)
	words += wc + wb
	images += ib + ic
	return words, images
}
