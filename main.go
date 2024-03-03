package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create a mock "database" to store customer data
type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type CustomerTable map[int]Customer

var customerTable CustomerTable = make(map[int]Customer)

func (c *CustomerTable) AddNewCustomer(customer Customer) error {

	// Customers in the database have unique ID values (i.e., no two customers have the same ID value).
	if _, inCustomerTable := customerTable[customer.ID]; inCustomerTable {
		fmt.Println("Customer with ID ", customer.ID, " already exists in the database.")
		return errors.New("Customer with ID " + strconv.Itoa(customer.ID) + " already exists in the database.")
	}
	customerTable[customer.ID] = customer
	return nil
}

func (c *CustomerTable) UpdateExistingCustomer(customer Customer) error {
	// Customers in the database have unique ID values (i.e., no two customers have the same ID value).
	if _, inCustomerTable := customerTable[customer.ID]; !inCustomerTable {
		fmt.Println("Customer with ID ", customer.ID, " does not exist in the database.")
		return errors.New("Customer with ID " + strconv.Itoa(customer.ID) + " does not exist in the database.")
	}
	customerTable[customer.ID] = customer
	return nil
}

func validateIdString(writer http.ResponseWriter, request *http.Request) (int, error) {
	vars := mux.Vars(request)
	idString := vars["id"]
	// convert id to int
	id, convertErr := strconv.Atoi(idString)

	if convertErr != nil {
		// writer.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(writer).Encode(nil)
		http.Error(writer, "Cannot convert id string \""+idString+"\" to int.", http.StatusNotFound)
		return 0, convertErr
	}

	return id, nil
}

// Getting a single customer through a /customers/{id} path
func getCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id, convertErr := validateIdString(writer, request)
	if convertErr != nil {
		return
	}

	if customer, inTable := customerTable[id]; inTable {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(customer)
	} else {
		// writer.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(writer).Encode(nil)
		http.Error(writer, "Customer with ID "+strconv.Itoa(id)+" does not exist in the database.", http.StatusNotFound)
	}
}

// Getting all customers through a the /customers path

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(customerTable)
}

// Creating a customer through a /customers path
func addCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var customer Customer
	_ = json.NewDecoder(request.Body).Decode(&customer)
	err := customerTable.AddNewCustomer(customer)

	if err != nil {
		// writer.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(writer).Encode(nil)
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(customerTable)
}

// Updating a customer through a /customers/{id} path
func updateCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id, convertErr := validateIdString(writer, request)
	if convertErr != nil {
		return
	}

	var customer Customer
	_ = json.NewDecoder(request.Body).Decode(&customer)

	if customer.ID != id {
		// writer.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(writer).Encode(nil)
		fmt.Println("Customer ID in the request body does not match the ID in the URL.")
		http.Error(writer, "Customer ID in the request body does not match the ID in the URL.", http.StatusBadRequest)
		return
	}

	err := customerTable.UpdateExistingCustomer(customer)

	if err != nil {
		// writer.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(writer).Encode(nil)
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(customer)
}

// Deleting a customer through a /customers/{id} path

func deleteCustomer(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	id, convertErr := validateIdString(writer, request)
	if convertErr != nil {
		return
	}

	if _, inTable := customerTable[id]; inTable {
		delete(customerTable, id)
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(customerTable)
	} else {
		// writer.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(writer).Encode(nil)
		http.Error(writer, "Customer with ID "+strconv.Itoa(id)+" does not exist in the database.", http.StatusNotFound)
	}
}

func main() {

	// Seed the database with initial customer data
	customers := []Customer{
		{ID: 1, Name: "John Doe", Role: "Admin", Email: "john.doe@fakejoxh.com", Phone: "123-456-7890", Contacted: false},
		{ID: 2, Name: "Jane Doe", Role: "User", Email: "jane.doe@fakejoxh.com", Phone: "123-456-7890", Contacted: false},
		{ID: 3, Name: "Jim Doe", Role: "User", Email: "jim.doe@fakejoxh.com", Phone: "123-456-7890", Contacted: false},
	}

	for _, customer := range customers {
		fmt.Println("Adding customer with ID ", customer.ID, " to the database.")
		customerTable.AddNewCustomer(customer)
	}

	// Create a new router
	router := mux.NewRouter()

	// Create a new route for getting a single customer
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(fs)

	port := ":3001"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}
