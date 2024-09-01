package main

import (
	"log"
	"net/http"

	"github.com/valentinrb1/go-api-rest.git/handlers"
)

func main() {

	usersMux := http.NewServeMux()

	usersMux.HandleFunc("/api/users/listall", handlers.GetUsersHandler)
	usersMux.HandleFunc("/api/users/createuser", handlers.CreateUserHandler)
	usersMux.HandleFunc("/api/users/login", handlers.LoginHandler)

	processingMux := http.NewServeMux()

	processingMux.HandleFunc("/api/processing/submit", handlers.SubmitProcessingHandler)
	processingMux.HandleFunc("/api/processing/summary", handlers.GetSummaryHandler)

	go func() {
		log.Println("Listening on port 8080 for users")
		log.Fatal(http.ListenAndServe(":8080", usersMux))
	}()

	log.Println("Listening on port 8081 for processing")
	log.Fatal(http.ListenAndServe(":8081", processingMux))
}
