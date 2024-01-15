package main

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

// Authorization middleware
func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your authorization logic goes here
		// For simplicity, let's check if a custom header "Authorization" is present in the request
		authHeader := r.Header.Get("Authorization")
		if authHeader != "mysecrettoken" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler if the authorization is successful
		next.ServeHTTP(w, r)
	})
}

// Handler for the protected resource
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a protected resource.")
}

func main() {
	r := mux.NewRouter()

	// Use the authorization middleware for this route
	r.Handle("/protected", authorize(http.HandlerFunc(protectedHandler))).Methods("GET")

	// Start the server
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
