package main
 
import (
    "fmt"
    "html/template"
    "net/http"
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
    http.HandleFunc("/", htmlHandler)
    http.ListenAndServe(":8080", nil)
}