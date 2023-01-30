package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandle)
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)
}
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word!!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
