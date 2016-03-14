package main

import (
	"net/http"
	"text/template"
	"github.com/nu7hatch/gouuid"
)

func startpage(res http.ResponseWriter, req *http.Request) {
	CreateCookie(res,req)
	var err error
	temp := template.Must(template.ParseFiles("index.html"))
	err = temp.Execute(res, nil)
	if err != nil {
		panic(err)
	}
}

func CreateCookie(res http.ResponseWriter,req *http.Request){
	cookie, err := req.Cookie("session-fino")
	if err == http.ErrNoCookie{
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String(),
			//Secure : true,
			HttpOnly:true,
		}
	}
	http.SetCookie(res,cookie)
}
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", startpage)
	http.ListenAndServe(":8080", nil)
}
