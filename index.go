package main

import (
	"net/http"
	"html/template"
	"log"
)

const counterTemplateStr = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>GO all things!</title>
			<style>
				body,html{height:100%;margin:0;font-family:sans-serif}.circle-holder{border-radius:50%;width:36px;height:36px;padding:8px;background:#fff;border:2px solid #666;color:#666;text-align:center;font-size:36px;position:relative;top:50%;-webkit-transform:translateY(-50%);-ms-transform:translateY(-50%);transform:translateY(-50%)}.full-page{width:100%;height:100%;min-height:100%;vertical-align:middle;text-align:-webkit-center}
			</style>
		</head>
		<body>
			<div class="full-page">
				<div class="circle-holder">
					{{.Counter}}
				</div>
			</div>
		</body>
	</html>
`
var counterTemplate = template.Must(template.New("template").Parse(counterTemplateStr))

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