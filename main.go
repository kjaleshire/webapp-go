package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/", HomeHandler)

    posts := r.Path("/posts").Subrouter()
    posts.Methods("GET").HandlerFunc(PostsIndexHander)
    posts.Methods("POST").HandlerFunc(PostsCreateHandler)

    post := r.PathPrefix("/posts/{id}").Subrouter()
    post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
    post.Methods("GET").HandlerFunc(PostShowHandler)
    post.Methods("PUT").HandlerFunc(PostUpdateHander)
    post.Methods("DELETE").HandlerFunc(PostDeleteHander)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("Starting server on", port)
    http.ListenAndServe(":" + port, r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Home")
}

func PostsIndexHander(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Posts create")
}
func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Posts edit")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    fmt.Fprintln(rw, "Posts show id", id)
}

func PostUpdateHander(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Post update")
}

func PostDeleteHander(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Post delete")
}
