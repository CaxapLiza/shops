package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/shops/services/outlet_employee/internal/handler"
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

	router.HandleFunc("/outlet_employees/list/{id}", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/outlet_employees/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/outlet_employees", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/outlet_employees/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/outlet_employees/{id}", handler.Delete).Methods("DELETE", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8006", router))
}
