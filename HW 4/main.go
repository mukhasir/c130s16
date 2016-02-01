package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", surfWebPage)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":8080", nil)
}

func surfWebPage(res http.ResponseWriter, req *http.Request) {

	var err error

	exersicePage := template.New("samplecode.html")
	exersicePage, err = exersicePage.ParseFiles("samplecode.html")

	if err != nil {
		log.Fatalln(err)
	}

	err = exersicePage.Execute(res, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
