package main

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
"google.golang.org/appengine/memcache"
	"fmt"
)

func storecookie(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("Test")
	if err == http.ErrNoCookie {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name : "Test",
			Value : id.String(),
			HttpOnly: true,
		}
	}
	http.SetCookie(res, cookie)
	fmt.Println(cookie.Value)
	setMemCache(res,req,cookie.Value)
	getMemCache(res,req)
}

func setMemCache(res http.ResponseWriter, req *http.Request,value string){
	ctx := appengine.NewContext(req)

	item1 := memcache.Item{
		Key:   "Test",
		Value: []byte(value),
	}

	memcache.Set(ctx, &item1)

	item, _ := memcache.Get(ctx, "Test")
	if item != nil {
		fmt.Fprintln(res, string(item.Value))
	}
}

func getMemCache(res http.ResponseWriter, req *http.Request){
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, "Test")
	fmt.Fprintln(res, item)
}

func main(){
	http.HandleFunc("/",storecookie)
	http.ListenAndServe(":8080",nil)
}
