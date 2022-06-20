package main

import (
	"fmt"
	"html/template"
	"log"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", htmlHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
