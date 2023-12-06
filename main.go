package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You did not authenticate"))
		return
	}
	w.Write([]byte("You sent this Authorization header: " + auth))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth", authHandler)
	mux.HandleFunc("/_/health", healthHandler)
	mux.HandleFunc("/_/readiness", readyHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server started on port 8080")
	log.Fatal(server.ListenAndServe())
}
