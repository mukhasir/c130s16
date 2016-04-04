package main

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
	"text/template"
)

func startpage(res http.ResponseWriter, req *http.Request) {
	var err error
	temp := template.Must(template.ParseFiles("index.html"))
	name := req.FormValue("name")
	age := req.FormValue("age")
	cookie, err := req.Cookie("session-fino")
	if err == http.ErrNoCookie {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			//Secure : true,
			HttpOnly: true,
		}
	} else {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + name + age,
			//Secure : true,
			HttpOnly: true,
		}
	}
	http.SetCookie(res, cookie)
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
