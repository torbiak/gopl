// ex7.9 serves an html table with a stable column sort.
//
// I used the Person type from the sort package's first example for ex7.8, so
// I'm also using it here instead of the Track type as specified in ex7.9.
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/torbiak/gopl/ex7.8"
)

var people = []column.Person{
	{"Alice", 20},
	{"Bob", 12},
	{"Bob", 20},
	{"Alice", 12},
}

var html = template.Must(template.New("people").Parse(`
<html>
<body>

<table>
	<tr>
		<th><a href="?sort=name">name</a></th>
		<th><a href="?sort=age">age</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Name}}</td>
		<td>{{.Age}}</td>
	</td>
{{end}}
</body>
</html>
`))

func main() {
	c := column.NewByColumns(people, 2)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "age":
			c.Select(c.LessAge)
		case "name":
			c.Select(c.LessName)
		}
		sort.Sort(c)
		err := html.Execute(w, people)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
