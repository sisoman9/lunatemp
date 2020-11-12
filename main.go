package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("site/*.gohtml"))
	http.HandleFunc("/", index)

	http.Handle("/includes/", http.StripPrefix("/includes", http.FileServer(http.Dir("./includes"))))

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.NotFound(w, req)
	}

}
