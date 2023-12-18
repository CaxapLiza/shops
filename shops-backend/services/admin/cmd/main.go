package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/shops/services/admin/internal/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	)

	router.Use(corsMiddleware)

	router.HandleFunc("/admins", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/admins/auth/{id}", handler.Authenticate).Methods("GET", "OPTIONS")
	router.HandleFunc("/admins/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/admins", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/admins/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/admins/{id}", handler.Delete).Methods("DELETE", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8003", router))
}
