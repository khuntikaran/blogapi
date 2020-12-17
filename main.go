package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to our website this is the private blog site where you can share your thout ")
}

func routehandler() {
	route := mux.NewRouter()
	route.HandleFunc("/", handler)
	route.HandleFunc("/blog", getAllBlog).Methods("GET")
	route.HandleFunc("/blog/{id}", getBlog).Methods("GET")
	route.HandleFunc("/blog/{id}", deleteBlog).Methods("DELETE")
	route.HandleFunc("/blog", createBlog).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", route))
}

func main() {
	routehandler()

}
