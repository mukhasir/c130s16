package main

import (
	"net/http"
	"text/template"
)

func startpage(res http.ResponseWriter, req *http.Request) {
	var err error
	temp := template.Must(template.ParseFiles("index.html"))
	err = temp.Execute(res, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", startpage)
	http.ListenAndServe(":8080", nil)
}
