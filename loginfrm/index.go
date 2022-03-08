package main 

import (
	"net/http"
	"fmt"
	// "github.com/gorilla/mux"
	// "html/template"
)

func main(){
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
	http.ListenAndServe(":8888",nil)
}

func index(w http.ResponseWriter, r *http.Request ){
	http.ServeFile(w,r,"login.html")
}

func login(w http.ResponseWriter, r *http.Request ){
	// http.ServeFile(w,r,"login.html")
	fmt.Println("Method:",r.Method)
	r.ParseForm()
	fmt.Println("Username:",r.Form["username"])
	fmt.Println("Password:",r.Form["password"])
}

