package main

<<<<<<< Updated upstream
import "fmt"
=======
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
}
>>>>>>> Stashed changes

func main(){
	fmt.Println("Hello, world!")
	fmt.Println("PR branch v2!")
}
