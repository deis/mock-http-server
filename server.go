package main

import (
	"fmt"
	"log"
	"net/http"
)

// Log the HTTP request
func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

// mockHandler responds with "ok" as the response body
func mockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

// rootHandler used to process all inbound HTTP requests
func rootHandler(w http.ResponseWriter, r *http.Request) {
	logHandler(w, r)
	mockHandler(w, r)
}

// Start an HTTP server which dispatches to the rootHandler
func main() {
	http.HandleFunc("/", rootHandler)
	port := "8080"
	log.Printf("server is listening on %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
