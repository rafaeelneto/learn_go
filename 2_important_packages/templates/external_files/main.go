package main

import (
	"html/template"
	"net/http"
	"os"
)

type TemplateData struct {
	Title   string
	Heading string
	Content string
	Items   []string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			"Meu titulo",
			"Heading",
			"Conteudozinho",
			[]string{"item1", "item2"},
		}

		// the name of the template must match the file name
		t := template.Must(template.New("index.html").ParseFiles("./templates/index.html"))

		w.Header().Add("Content-Type", "application/xml")
		err := t.Execute(w, data)

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func onlyPrint() {
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
