package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/shops/services/employee/internal/handler"
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

	router.HandleFunc("/employees", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/employees/auth/{id}", handler.Authenticate).Methods("GET", "OPTIONS")
	router.HandleFunc("/employees/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/employees", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/employees/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/employees/{id}", handler.Delete).Methods("DELETE", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8005", router))
}
