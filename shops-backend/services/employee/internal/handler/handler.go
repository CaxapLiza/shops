package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/student/shops/services/common"
	"github.com/student/shops/services/employee/internal"
	"github.com/student/shops/services/employee/internal/repository"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func GetList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	employees, err := repo.GetList()
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	admins, err := repo.Authenticate(id)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(admins)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Get(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	employees, err := repo.Get(id)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var requestBody struct {
		Name      string `json:"name"`
		Itn       string `json:"itn"`
		Passport  string `json:"passport"`
		Snils     string `json:"snils"`
		Phone     string `json:"phone"`
		AccountId int    `json:"account_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	employee := &internal.Employee{Name: requestBody.Name, Itn: requestBody.Itn, Passport: requestBody.Passport, Snils: requestBody.Snils, Phone: requestBody.Phone, AccountId: requestBody.AccountId}

	if err := repo.Create(employee); err != nil {
		http.Error(w, "Error creating", http.StatusInternalServerError)
		log.Println("Create Error:", err)
		return
	}

	response, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	var requestBody struct {
		Name      string `json:"name"`
		Itn       string `json:"itn"`
		Passport  string `json:"passport"`
		Snils     string `json:"snils"`
		Phone     string `json:"phone"`
		AccountId int    `json:"account_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	if err := repo.Update(id, requestBody.Name, requestBody.Itn, requestBody.Passport, requestBody.Snils, requestBody.Phone, requestBody.AccountId); err != nil {
		http.Error(w, "Error updating", http.StatusInternalServerError)
		log.Println("Update Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	if err := repo.Delete(id); err != nil {
		http.Error(w, "Error deleting", http.StatusInternalServerError)
		log.Println("Delete Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
