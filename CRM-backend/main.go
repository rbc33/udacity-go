package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

const jsonFile = "data/data.json"

// "gET"AllUsers devuelve todos los usuarios desde el archivo JSON.
func getAllCustomers() ([]Customer, error) {
	var users []Customer
	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &users)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

// saveUsers guarda los usuarios en el archivo JSON.
func saveCustomer(customer []Customer) error {
	data, err := json.MarshalIndent(customer, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonFile, data, 0644)
}

// GetUsersHandler maneja las solicitudes GET /users.
func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	users, err := getAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// AddUserHandler maneja las solicitudes POST /users.
func AddCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	customers, err := getAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer.ID = strconv.Itoa(len(customers) + 1) // Generar un nuevo ID
	customers = append(customers, customer)

	if err := saveCustomer(customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}
func updateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]

	customers, err := getAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer.ID = id
	customers = append(customers, customer)

	if err := saveCustomer(customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	customers, err := getAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, c := range customers {
		if c.ID == id {
			customer := c

			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(customer)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

	}
}
func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	customers, err := getAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	found := false
	for i, u := range customers {
		if u.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := saveCustomer(customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(customers)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customer", GetCustomerHandler).Methods("GET")
	r.HandleFunc("/customer", AddCustomerHandler).Methods("POST")
	r.HandleFunc("/customer/{id}", getCustomerHandler).Methods("GET")
	r.HandleFunc("/customer/{id}", DeleteCustomerHandler).Methods("DELETE")
	r.HandleFunc("/customer/{id}", updateCustomerHandler).Methods("PUT")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", r))
}
