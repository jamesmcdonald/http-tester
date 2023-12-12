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

func logWrap(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/auth", logWrap(authHandler))
	mux.Handle("/_/health", logWrap(healthHandler))
	mux.Handle("/_/readiness", logWrap(readyHandler))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server started on port 8080")
	log.Fatal(server.ListenAndServe())
}
