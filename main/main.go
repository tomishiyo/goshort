package main

import (
	"fmt"
	"net/http"
)

// One of the contents to be served
func main_page(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

// This function writes the Response Headers
func write_headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Links the handlers to the functions
	http.HandleFunc("/", main_page)
	http.HandleFunc("/headers", write_headers)

	http.ListenAndServe("0.0.0.0:80", nil)
}
