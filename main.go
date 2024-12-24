package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", handleReadiness)
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	log.Println("Listening on port :8080")
	log.Fatal(server.ListenAndServe())
}

func handleReadiness(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
