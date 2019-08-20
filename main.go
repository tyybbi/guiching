package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pageVars struct {
	PId    int
	RId    int
	PName  string
	RName  string
	PDesc  string
	RDesc  string
	PLines [6]string
	RLines [6]string
	Rel    bool
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", createShapes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(w http.ResponseWriter, tmpl string, pageVars pageVars) {

	tmpl = fmt.Sprintf("template/%s", tmpl)
	t, err := template.ParseFiles(tmpl)

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, pageVars)

	if err != nil {
		log.Print("template executing error: ", err)
	}
}
