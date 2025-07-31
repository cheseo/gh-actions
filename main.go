package main

import (
	"fmt"
	"net/http"
	"log"
)

func add(a, b int) int {
	return a+b
}

func handler(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	if name == "" {
		name = "unknown"
	}
	fmt.Fprintf(w, "Hello %s!!\n", name)
	fmt.Fprintf(w, "help: use this:")
	fmt.Fprintf(w, "use /?name=name to display your name!")
}

func main(){
	http.HandleFunc("/", handler)
	log.Println("Listening on 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
