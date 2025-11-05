package main

import (
	"fmt"
	"gopl/test"
	"html/template"
	"log"
	"net/http"
)

var issueListForServer = template.Must(template.New("issuelist").Parse(`
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
	http.HandleFunc("/issues", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	result, err := test.SearchIssues([]string{"repo:golang/go", "commenter:gopherbot", "json", "encoder"})
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	if err = issueListForServer.Execute(w, result); err != nil {
		fmt.Fprintln(w, err)
		return
	}
}
