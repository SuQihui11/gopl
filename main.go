package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	//var w test.Wheel
	//// 此时即使point和circle不导出，我们仍然可以使用简短形式来访问w.X
	//// 但是w.circle.point.X = 8会出错
	//w.X = 8
	//w.Y = 8
	//w.Radius = 5
	//w.Spokes = 20
	//
	//fmt.Printf("%#v\n", w)

	//data, err := json.MarshalIndent(test.Movies, "", "   ")
	//if err != nil {
	//	log.Fatalf("JSON marshaling failed: %s", err)
	//}
	//fmt.Printf("%s\n", data)
	//
	//var titles []struct{ Title string }
	//if err := json.Unmarshal(data, &titles); err != nil {
	//	log.Fatalf("JSON unmarshaling failed: %s", err)
	//}
	//fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

	//result, err := test.SearchIssues(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
	//
	//hash := test.ProcessData(result)
	//for key, val := range hash {
	//	fmt.Printf("%s: %s\n", key, val)
	//}

	//test.GetXKCDAndSave()
	//key := os.Args[1:]
	//input := strings.Join(key, "")
	//fmt.Println("输入参数：", input)
	//fmt.Println(test.ResImgUrl(key[0]))

	//result, err := test.SearchIssues(os.Args[1:])
	//if err != nil {
	//	fmt.Println(err)
	//}
	////if err := report.Execute(os.Stdout, result); err != nil {
	////	fmt.Println(err)
	////}
	//if err := issueList.Execute(os.Stdout, result); err != nil {
	//	fmt.Println(err)
	//}

	//可以通过对信任的HTML字符串使用template.HTML类型来抑制这种自动转义的行为
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
