package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Print the URL
    fmt.Printf("Received %s request for URL: %s\n", r.Method, r.URL.Path)
    
    // Write response
    w.Header().Set("Content-Type", "text/html")
    if r.Method == "GET" {
        fmt.Fprintln(w, "Hello, world!")
    } else if r.Method == "POST" {
        fmt.Fprintln(w, "POST request received!")
    }
}

func main() {
    http.HandleFunc("/", handler)
    port := 8080
    fmt.Printf("Starting server on port %d\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

