// #nosec
package main

import (
	"html/template"
	"log"
	"os"
)

type Location struct {
	Lat float64
	Lng float64
}

var tmplText = `
I am at:
{{ .Lat }}/{{ .Lng }}
`

func main() {
	tmpl, err := template.New("loc").Parse(tmplText)
	if err != nil {
		log.Fatal(err)
	}

	loc := Location{1.2, 3.4}
	tmpl.Execute(os.Stdout, loc)
}
