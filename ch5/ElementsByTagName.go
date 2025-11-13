package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		return nil
	}

	var result []*html.Node

	// 创建set
	nameSet := make(map[string]bool)
	for _, name := range name {
		nameSet[name] = true
	}

	// 匿名函数可以直接去访问函数外部的局部变量，减少参数的传递
	var traverse func(n *html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && nameSet[n.Data] {
			result = append(result, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)
	return result
}

func parseHTML(text string) (n *html.Node, err error) {
	n, err = html.Parse(strings.NewReader(text))
	return
}

// 辅助函数：打印节点信息
func printNode(n *html.Node) {
	if n.Type == html.ElementNode {
		attrs := []string{}
		for _, attr := range n.Attr {
			attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
		}
		attrStr := ""
		if len(attrs) > 0 {
			attrStr = " " + strings.Join(attrs, " ")
		}
		fmt.Printf("<%s%s>\n", n.Data, attrStr)
	}
}

func main() { // 示例HTML内容
	htmlContent := `
	<html>
	<head>
		<title>测试页面</title>
		<meta charset="utf-8">
	</head>
	<body>
		<h1>主标题</h1>
		<p>第一段</p>
		<p>第二段</p>
		<div class="container">
			<h2>子标题</h2>
			<p>第三段</p>
			<span>内联文本</span>
			<img src="photo.jpg" alt="示例图片">
		</div>
		<ul>
			<li>列表项1</li>
			<li>列表项2</li>
		</ul>
	</body>
	</html>
	`
	doc, err := parseHTML(htmlContent)
	if err != nil {
		fmt.Println(err)
	}
	// 示例1：查找所有 <p> 和 <h1> 元素
	fmt.Println("=== 示例1：查找 <p> 和 <h1> ===")
	elems1 := ElementsByTagName(doc, "p", "h1")
	for _, elem := range elems1 {
		printNode(elem)
	}
	// 输出:
	// <h1>
	// <p>
	// <p>
	// <p>

	// 示例2：查找所有 <meta> 和 <img> 元素
	fmt.Println("\n=== 示例2：查找 <meta> 和 <img> ===")
	elems2 := ElementsByTagName(doc, "meta", "img")
	for _, elem := range elems2 {
		printNode(elem)
	}
	// 输出:
	// <meta charset="utf-8">
	// <img src="photo.jpg" alt="示例图片">

	// 示例3：查找单个元素（兼容单标签查询）
	fmt.Println("\n=== 示例3：查找 <title> ===")
	elems3 := ElementsByTagName(doc, "title")
	for _, elem := range elems3 {
		printNode(elem)
	}
	// 输出:
	// <title>
}
