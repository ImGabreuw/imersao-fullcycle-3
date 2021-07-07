package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8000", nil)
}

func HomePage(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./resources/index.html"))
	t.Execute(writer, nil)
}
