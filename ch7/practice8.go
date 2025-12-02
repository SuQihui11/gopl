//package main
//
//import (
//	"fmt"
//	"os"
//	"sort"
//	"text/tabwriter"
//	"time"
//)
//
//// Track 是我们要排序的数据模型
//type Table struct {
//	Title  string
//	Artist string
//	Album  string
//	Year   int
//	Length time.Duration
//}
//
//// 模拟一些初始数据
//var tables = []*Table{
//	{"Go", "Delilah", "From the Roots Up", 2012, lengthFromStr("3m38s")},
//	{"Go", "Moby", "Moby", 1992, lengthFromStr("3m37s")},
//	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, lengthFromStr("4m24s")},
//	{"Go Ahead", "Alicia Keys", "As I Am", 2007, lengthFromStr("4m36s")},
//}
//
//func lengthFromStr(s string) time.Duration {
//	d, err := time.ParseDuration(s)
//	if err != nil {
//		panic(s)
//	}
//	return d
//}
//
//func printTables(tables []*Table) {
//	const format = "%v\t%v\t%v\t%v\t%v\t\n"
//	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
//	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
//	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
//	for _, t := range tables {
//		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
//	}
//	tw.Flush()
//}
//
//type LessFunc func(a, b *Table) bool
//
//type multiSort struct {
//	tables    []*Table
//	lessFuncs []LessFunc
//}
//
//func (m *multiSort) Len() int {
//	return len(m.tables)
//}
//
//func (m *multiSort) Swap(i, j int) {
//	m.tables[i], m.tables[j] = m.tables[j], m.tables[i]
//}
//
//func (m *multiSort) Less(i, j int) bool {
//	p, q := m.tables[i], m.tables[j]
//	for _, less := range m.lessFuncs {
//		if less(p, q) {
//			return true
//		}
//		if less(q, p) {
//			return false
//		}
//	}
//	return false
//}
//
//func (m *multiSort) Select(less LessFunc) {
//	m.lessFuncs = append([]LessFunc{less}, m.lessFuncs...)
//	if len(m.lessFuncs) > 5 {
//		m.lessFuncs = m.lessFuncs[:5]
//	}
//}
//
//func main() {
//	ms := &multiSort{tables: tables}
//	fmt.Println("\n--- 1. 初始状态 (未排序) ---")
//	printTables(tables)
//
//	fmt.Println("\n--- 2. 点击 'Length' ---")
//	ms.Select(func(a, b *Table) bool {
//		return a.Length < b.Length
//	})
//	sort.Sort(ms)
//	printTables(ms.tables)
//
//	// 模拟用户点击：按年份(Year)排序
//	// Year 成为主键，Length 成为次键
//	fmt.Println("\n--- 3. 点击 'Year' ---")
//	ms.Select(func(x, y *Table) bool {
//		return x.Year < y.Year
//	})
//	sort.Sort(ms)
//	printTables(tables)
//
//	// 模拟用户点击：按标题(Title)排序
//	// Title 成为主键，Year 成为次键，Length 成为第三键
//	// 注意观察：Title 相同的行，会保持 Year 的顺序
//	fmt.Println("\n--- 4. 点击 'Title' ---")
//	ms.Select(func(x, y *Table) bool {
//		return x.Title < y.Title
//	})
//	sort.Sort(ms)
//	printTables(tables)
//}
