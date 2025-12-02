package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track 是我们要排序的数据模型
type Table struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// 模拟一些初始数据
var tables = []*Table{
	{"Go", "Delilah", "From the Roots Up", 2012, lengthFromStr("3m38s")},
	{"Go", "Moby", "Moby", 1992, lengthFromStr("3m37s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, lengthFromStr("4m24s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, lengthFromStr("4m36s")},
}

func lengthFromStr(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTables(tables []*Table) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tables {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type LessFunc func(a, b *Table) bool

type multiSort struct {
	tables    []*Table
	lessFuncs []LessFunc
}

func (m *multiSort) Len() int {
	return len(m.tables)
}

func (m *multiSort) Swap(i, j int) {
	m.tables[i], m.tables[j] = m.tables[j], m.tables[i]
}

func (m *multiSort) Less(i, j int) bool {
	p, q := m.tables[i], m.tables[j]
	for _, less := range m.lessFuncs {
		if less(p, q) {
			return true
		}
		if less(q, p) {
			return false
		}
	}
	return false
}

func (m *multiSort) Select(less LessFunc) {
	m.lessFuncs = append([]LessFunc{less}, m.lessFuncs...)
	if len(m.lessFuncs) > 5 {
		m.lessFuncs = m.lessFuncs[:5]
	}
}

// 初始化全局的排序器
var ms = &multiSort{tables: tables}

var trackTable = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html>
<head>
<style>
table { border-collapse: collapse; width: 50%; }
th, td { text-align: left; padding: 8px; border-bottom: 1px solid #ddd; }
th { background-color: #f2f2f2; }
a { text-decoration: none; color: black; font-weight: bold; }
a:hover { color: blue; }
</style>
</head>
<body>

<h2>Music Tracks</h2>
<table>
  <tr>
    <th><a href="/?sort=Title">Title</a></th>
    <th><a href="/?sort=Artist">Artist</a></th>
    <th><a href="/?sort=Album">Album</a></th>
    <th><a href="/?sort=Year">Year</a></th>
    <th><a href="/?sort=Length">Length</a></th>
  </tr>
  {{range .}}
  <tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
  </tr>
  {{end}}
</table>

</body>
</html>
`))

func handleTracks(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("sort")
	// 根据参数更新排序器的状态
	switch key {
	case "Title":
		ms.Select(func(a, b *Table) bool { return a.Title < b.Title })
	case "Artist":
		ms.Select(func(a, b *Table) bool { return a.Artist < b.Artist })
	case "Album":
		ms.Select(func(a, b *Table) bool { return a.Album < b.Album })
	case "Year":
		ms.Select(func(a, b *Table) bool { return a.Year < b.Year })
	case "Length":
		ms.Select(func(a, b *Table) bool { return a.Length < b.Length })
	}

	// 执行排序
	// 注意：因为 ms 保存了 lessFuncs 的历史，所以它是"有状态"的。
	// 如果你先点 Year 再点 Title，Year 会变成次要排序键。
	sort.Sort(ms)

	// 渲染 HTML
	if err := trackTable.Execute(w, ms.tables); err != nil {
		log.Printf("template execution failed: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handleTracks)
	log.Println("Listening on http://localhost:8000 ...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
