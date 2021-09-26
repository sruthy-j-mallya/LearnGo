package main

import (
	"net/http"
	"html/template"
	"log"
)

var tmpl *template.Template

type Todo struct{
	Item string
	Done bool
}

type PageData struct{
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
		},
	}
	tmpl.Execute(w,data)
}

func main(){
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("/templates/index.html"))
	mux.HandleFunc("/todo",todo)
	log.Fatal(http.ListenAndServe(":8080",mux))
}