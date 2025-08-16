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
	tmp := template.Must(template.New("base").ParseGlob("./templates/*.html"))

	err := tmp.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
