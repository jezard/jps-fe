package main

import (
	"fmt"
	"net/http"
	//"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello world")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello home")
}
