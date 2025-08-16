package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// path based on where the command is running
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)

	// can work with multiple routes
	mux.HandleFunc("/api/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("my blog content"))
	})

	http.ListenAndServe(":8080", mux)
}
