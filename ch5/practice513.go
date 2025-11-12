package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func Extract513(rawURL string) ([]string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", rawURL, err)
	}

	var urls []string

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				urls = append(urls, link.String())
			}
		}
	}

	forEach(doc, visitNode, nil)
	return urls, nil
}

func forEach(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEach(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawlAndSave(rawURL string, baseDomain string) []string {
	fmt.Printf("crawling %s\n", rawURL)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Printf("Error parsing URL: %s\n", rawURL)
		return nil
	}

	if parsedURL.Host != baseDomain {
		fmt.Printf("URL does not belong to base domain: %s\n", parsedURL)
		return nil
	}

	// http 请求
	resp, err := http.Get(rawURL)
	if err != nil {
		log.Printf("Error fetching URL: %s\n", rawURL)
		return nil
	}
	// 状态检查
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching URL: %s\n", rawURL)
		return nil
	}

	// 保存页面
	if err := savePage513(rawURL, resp.Body, "my_crawled_pages"); err != nil {
		log.Printf("save page error: %v", err)
	}
	// 提取页面的链接
	links, err := Extract513(rawURL)
	if err != nil {
		log.Printf("Error extracting URL: %s\n", rawURL)
	}
	return links
}

func savePage513(rawURL string, body io.Reader, outputDir string) error {
	// 正确的解析这里的url然后创建对应的目录
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	path := parsedURL.Path
	if path == "/" {
		path = "/index.html"
	}
	// 安全化文件名
	safePath := strings.TrimPrefix(path, "/")
	safePath = strings.ReplaceAll(safePath, "/", string(filepath.Separator))

	fullPath := filepath.Join(outputDir, safePath)

	// 创建所有必要目录
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("create file error: %v", err)
	}
	defer file.Close()
	if _, err := io.Copy(file, body); err != nil {
		return fmt.Errorf("copy error: %v", err)
	}
	return nil
}

// breadthFirst遍历
func breadthFirst513(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		// 逐个处理当前层
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// 追加每一个当前层元素的下一层节点，然后处理
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func main() {
	startUrl := "https://www.kimi.com/"
	parsedUrl, err := url.Parse(startUrl)
	if err != nil {
		log.Fatal(err)
	}
	baseDomain := parsedUrl.Hostname()
	fmt.Printf("baseDomain: %s\n", baseDomain)
	// 使用闭包捕获baseDomain
	crawl := func(url string) []string {
		return crawlAndSave(url, baseDomain)
	}

	breadthFirst513(crawl, []string{startUrl})
}
