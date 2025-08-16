package main

import "net/http"

type blog struct {
	content string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.content))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hollo"))
	})
	mux.HandleFunc("/subpath", subpathHandle)
	mux.Handle("/blog", blog{content: "aaaaa"})

	// ListerAndServe is a blocking function
	go http.ListenAndServe(":8080", mux)

	muxSecond := http.NewServeMux()
	muxSecond.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hollo rafael"))
	})
	// ListerAndServe is a blocking function
	http.ListenAndServe(":8085", muxSecond)
}

func subpathHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("subtreee holo"))
}
