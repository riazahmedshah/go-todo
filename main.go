package main

import (
	"log"
	"net/http"

	"github.com/riazahmedshah/todo/routers"
	"github.com/riazahmedshah/todo/stores"
)

func main() {
	if err := stores.InitDB(); err != nil {
		log.Fatalf("Failed to Connect DataBase: %v", err)
	}
	mux := http.NewServeMux()

	routers.SetupRoutes(mux)

	const port = ":3000"
	addr := "http://localhost" + port

	log.Printf("Server started on %s", addr)
	log.Println("Press Ctrl+C to stop the server...")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
