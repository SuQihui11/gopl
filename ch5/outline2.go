package main

import (
	"golang.org/x/net/html"
)

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func expand(s string, f func(string) string) string {
	// 直接使用现有的工具
	//s = strings.ReplaceAll(s, "foo", f("foo"))
	//return s

	return helper(s, 0, len(s), f)
}

func helper(s string, start, end int, f func(string) string) string {
	// 小于直接返回
	if end-start < 3 {
		return s
	}

	// 如果是目标str，直接进行处理然后继续处理后续的部分即t
	if s[start:start+3] == "foo" {
		t := s[start+3 : end]
		return f(s[start:start+3]) + helper(t, 0, len(t), f)
	}

	// 如果当前的str不符合目标条件，那么就前进一位进行处理
	t := s[start+1 : end]
	return s[start:start+1] + helper(t, 0, len(t), f)
}
