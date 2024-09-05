package main

import "net/http"
import "fmt"
import "io/ioutil"
import "os"

func main(){
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME")
	version := os.Getenv("VERSION")
	fmt.Fprintf(w, "Hello %s! The current version is %s!", name, version)
}

func ConfigMap(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile("/app/message/message.txt")
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	
	fmt.Fprintf(w, "message.txt: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request){
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s, Password: %s", user, password)
}