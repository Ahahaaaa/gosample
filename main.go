package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	
)


func main(){
	fmt.Println("start")
	http.HandleFunc("/",handler)
	port:=os.Getenv("PORT")
	if port ==""{
		port="8000"
	}
	err:=http.ListenAndServe(":"+port,nil)
	if err !=nil{
		log.Fatal(err)
	}
}
func handler(w http.ResponseWriter,r*http.Request){
	fmt.Fprintf(w,"<h1>httpserverへようこそ</h1>\n")
	if err:=r.ParseForm();err!=nil{
		log.Print(err)
	}
}