package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(templateText string, data interface{}) {
	tmpl, err := template.New("template").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	// YOUR CODE HERE:
	// Set templateText so that the calls to executeTemplate
	// below will produce the output shown.
	templateText := "{{range .}}- {{.}}\n{{end}}"
	executeTemplate(templateText,
		[]string{"apples", "oranges", "pears"})
	executeTemplate(templateText,
		[]string{"chicken", "beef", "lamb"})
}