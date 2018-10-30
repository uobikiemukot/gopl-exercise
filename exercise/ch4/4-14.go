package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopl.io/ch4/github"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<form action="/" method="get">
label:     <input type="text" name="label" size="20">
milestone: <input type="text" name="milestone" size="20">
author:    <input type="text" name="author" size="20">
assignee:  <input type="text" name="assignee" size="20">
<input type="submit" value="submit">
</form>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Author</th>
  <th>Title</th>
  <th>Age</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.CreatedAt | daysAgo}} days</td>
</tr>
{{end}}
</table>
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ./rep_info user/repo\n")
		os.Exit(1)
	}

	// search issues byt label, milestone, assignee and author
	handler := func(w http.ResponseWriter, r *http.Request) {
		repo := os.Args[1]
		report, err := template.New("report").
			Funcs(template.FuncMap{"daysAgo": daysAgo}).
			Parse(templ)
		if err != nil {
			log.Fatal(err)
		}

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		var q = []string{fmt.Sprintf("repo:%s", repo)}
		// r.Form ==> url.Values ==> map[string][]string
		for k, v := range r.Form {
			if len(v) == 0 {
				continue
			}

			s := strings.Join(v, " ")
			if len(s) == 0 {
				continue
			}

			switch k {
			case "label":
				q = append(q, fmt.Sprintf("label:%s", s))
			case "milestone":
				q = append(q, fmt.Sprintf("milestone:%s", s))
			case "assignee":
				q = append(q, fmt.Sprintf("assignee:%s", s))
			case "author":
				q = append(q, fmt.Sprintf("author:%s", s))
			default:
				fmt.Fprintf(os.Stderr, "unknown key/values: key:%s / values:%v\n", k, v)
			}
		}

		fmt.Fprintf(os.Stderr, "q:\n%s\n", strings.Join(q, "\n"))
		result, err := github.SearchIssues(q)
		if err != nil {
			log.Fatal(err)
		}

		if err := report.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
