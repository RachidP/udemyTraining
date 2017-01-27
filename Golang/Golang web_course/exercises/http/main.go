package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	mux.HandleFunc("/bar", bar)
	mux.HandleFunc("/about", about)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {

		fmt.Println("errore ListeeAnd Serve", err)

	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "foo.gohtml", req.Method)

}

func bar(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "bar.gohtml", req.URL.Path)
	//tpl.ExecuteTemplate(w, "foo.gohtml", nil)
}

func about(w http.ResponseWriter, req *http.Request) {
	d := struct{ Fname, Lname string }{"Pippo", "Pluto"}
	tpl.ExecuteTemplate(w, "about.gohtml", d)
}
