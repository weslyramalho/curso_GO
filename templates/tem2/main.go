package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Curso struct {
	Nome          string
	CaragaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	/*
		t:= template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper()})
		t=template.Must(t.ParseFiles(templates...))
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(os.Stdout, Cursos{
			{"Go", 40},
			{"Java", 20},
			{"phyton", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)
}
