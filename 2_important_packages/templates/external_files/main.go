package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	Title   string
	Heading string
	Content string
	Items   []string
}

func main() {
	data := TemplateData{
		"Meu titulo",
		"Heading",
		"Conteudozinho",
		[]string{"item1", "item2"},
	}

	// the name of the template must match the file name
	t := template.Must(template.New("index.html").ParseFiles("./templates/index.html"))

	err := t.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
