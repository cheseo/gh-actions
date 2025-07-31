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
	fmt.Fprintf(w, "Hello %d!!\n", add(2, 3))
}

func main(){
	http.HandleFunc("/", handler)
	log.Println("Listening on 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
