package main

import (
	"fmt"
	"log"
	"net/http"
)

// func helloHandler(res *http.ResponseWriter,req *http.Request){
// 	if res.u
// }
func main() {
	fileserver:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",func(w http.ResponseWriter, r *http.Request) {
		if r.Method=="POST" {
			if err:=r.ParseForm(); err!=nil {
				fmt.Print(err);
				return
			}
			name:=r.FormValue("username");
			fmt.Print(name);
		}
	})
	http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		if r.Method!="GET" {
			http.Error(w,"method is not supported",http.StatusNotFound);
			return
		}
		fmt.Print(w,"Hello");
	})
	fmt.Print("starting server at port 3000")
	err:=http.ListenAndServe("127.0.0.1:3000",nil)
	if err!=nil {
		log.Fatal(err)
	}
}