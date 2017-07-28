package main

import (
	"os"
	"fmt"
	"net/http"
	"html/template"
	"log"
	"github.com/go-redis/redis"
)

type CounterPage struct {
	Host string
	Counter string
}

const counterTemplateStr = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>GO all things!</title>
			<script>
					function init() {
						var color ='#'+Math.random().toString(16).substr(2,6);
						var holder = document.getElementsByClassName("counter-holder")[0];
						holder.style.color = color;
						holder.style.borderColor = color;			
					}
			</script>
			<style>
				body,html{height:100%;margin:0;font-family:Impact}.counter-holder{padding:8px;text-align:center;font-size:104px;position:relative;top:50%;-webkit-transform:translateY(-50%);-ms-transform:translateY(-50%);transform:translateY(-50%)}.full-page{width:100%;height:100%;min-height:100%;vertical-align:middle;text-align:-webkit-center}.hostname{position:relative;top:50%;}    
			</style>
		</head>
		<body onload="init();">
			<div class="full-page">
				<div class="counter-holder">
					{{.Counter}}
				</div>
				<div class="hostname">
					~{ {{.Host}} }~
				</div>
			</div>
		</body>
	</html>
`

var counterTemplate = template.Must(template.New("template").Parse(counterTemplateStr))

var client = redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

func counterHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal("error getting hostname")
		panic(err)
	}
	val := getAndIncrement();
	data := &CounterPage{ 
		Counter : val,
		Host : name,
	}
	if err := counterTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

func getAndIncrement() string {
	if err := client.Incr("counter").Err(); err != nil {
		panic(err)
	}
	val, err := client.Get("counter").Int64()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", val)
}

func main() {
	if err := client.Set("counter", 0, 0).Err(); err != nil {
		panic(err)
	}
	http.HandleFunc("/att", counterHandler)
	log.Println("server up and running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}