package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to my server")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello %s\n", name)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/", loggingMiddleware(homeHandler))
	http.HandleFunc("/hello", loggingMiddleware(greetHandler))
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	fmt.Println("Server is listening on port 8080:")
	http.ListenAndServe(":8080", nil)
}
