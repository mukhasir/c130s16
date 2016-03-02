package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init(){
	var err error
	tpl, err = template.ParseFiles("templates/index.gohtml")
	//tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	if err != nil {
		//log.Fatalln(err)
		panic(err)
	}
}

func main(){
	http.HandleFunc("/", index)
	//http.Handle("/css/",http.StripPrefix("/css",http.FileServer(http.Dir("public/css"))))
	//http.Handle("/pic/",http.StripPrefix("/pic",http.FileServer(http.Dir("public/pic"))))
	http.Handle("/public/",http.StripPrefix("/public",http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080",nil)
}

func index(res http.ResponseWriter, req *http.Request){
	tpl.Execute(res,nil)
}
