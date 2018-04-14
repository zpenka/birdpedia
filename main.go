package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func newRouter() *mux.Router {
  r:= mux.NewRouter()
  r.HandleFunc("/hello", handler).Methods("GET")
  return r
}

func main() {
  // Declare router
  r := newRouter()
  http.ListenAndServe(":8080", r)
}

func handler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!")
}
