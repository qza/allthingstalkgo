package main

import (
	"net/http"
	"html/template"
	"log"
)

var counterTemplate = template.Must(template.ParseFiles("index.tmpl"))

func main() {
	http.HandleFunc("/", counterHandler)
	log.Println("server up and running")
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

type CounterPage struct {
	Counter string
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	data := &CounterPage{ 
		Counter : "1",
	}
	if err := counterTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}