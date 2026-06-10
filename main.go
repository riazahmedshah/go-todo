package main

import (
	"log"
	"net/http"
)

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("Failed to Connect DataBase: %v", err)
	}
	mux := http.NewServeMux()

	setupRoutes(mux)

	const port = ":3000"
	addr := "http://localhost" + port

	log.Printf("Server started on %s", addr)
	log.Println("Press Ctrl+C to stop the server...")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
