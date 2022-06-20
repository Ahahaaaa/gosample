package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)
 
func htmlHandler(w http.ResponseWriter, r *http.Request) {
 
    tmp, _ := template.ParseFiles("tem.html")
 
    values := map[string]string{
        "account": "user-1",
        "name":    "さとうこうや",
        "passwd":  "sample-pass",
    }
 
    if err := tmp.Execute(w, values); err != nil {
        fmt.Println(err)
    }
}
 

func main() {
    port:=os.Getenv("PORT")
    if  port==""{
    } else {
        port = "8080"
    }
    http.HandleFunc("/", htmlHandler)
    http.ListenAndServe(":"+port, nil)
}