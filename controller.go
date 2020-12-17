package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllBlog(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blog)
}

func getBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, blogs := range blog {
		if blogs.ID == key {
			json.NewEncoder(w).Encode(blogs)
		}
	}
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for index, blogs := range blog {
		if blogs.ID == key {
			blog = append(blog[:index], blog[index+1:]...)
			json.NewEncoder(w).Encode(blog)
		}
	}
}

func createBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBlog Blog
	json.NewDecoder(r.Body).Decode(&newBlog)
	newBlog.ID = strconv.Itoa(len(blog) + 1)
	blog = append(blog, newBlog)
	json.NewEncoder(w).Encode(newBlog)

}
