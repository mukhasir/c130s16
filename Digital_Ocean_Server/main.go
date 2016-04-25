package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request) {
		io.WriteString(res,"Man of my words")
	})
	http.ListenAndServe(":9000",nil)
}
