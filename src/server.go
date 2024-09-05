package main

import "net/http"
import "fmt"
import "os"

func main(){
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME")
	version := os.Getenv("VERSION")
	fmt.Fprintf(w, "Hello %s! The current version is %s!", name, version)
}