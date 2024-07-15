package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handler)
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        fmt.Println(err)
    }
}
