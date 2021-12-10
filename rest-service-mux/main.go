package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee Struct type(Model)
type Employee struct{
	ID string `json:"ID"`
	FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    Salary  float32 `json:"salary"`
	Address *Address `json:"address"`
}

// Address Struy type(Model)
type Address struct{
	StreetName string `json:"streetname"`
	PinCode int `json:"pincode"`
	State string `json:"state"`
	Country string `json:"country"`
}

// EmployeeResponse Struct type
type EmployeeResponse struct{
    Message string `json:"message"`
}

// Initialize employess var as slice Employee struct
var employees []Employee

// Get All Employees
func getEmployees(w http.ResponseWriter, r *http.Request){
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(employees)
}

// Get Employee by Id
func getEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	// Get request parameters
	params := mux.Vars(r)
	log.Println("Request params:", params)
    // Loop through employees to  find id
	for _, item := range employees {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
    json.NewEncoder(w).Encode(&EmployeeResponse{Message: "No employee record found with Id:"+params["id"]})
}


// Create New Employee
func createEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = strconv.Itoa(rand.Intn(10000000))  // Mock ID - not safe
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
	
}

// // Update an employee
// func updateEmployee(w http.ResponseWriter, r *http.Request)  {
	
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars()
// }

func main() {
	log.Println("Welcome to Rest API Demo using Mux")
 
	// Mock data @TODO - Implement database
	employees = append(employees, Employee{ID: "101", FirstName: "Sangram", LastName: "Konde", Salary: 2000.00, Address: &Address{
		StreetName: "Pune", PinCode: 401058, State: "Maharashtra", Country: "India",
	}})
	employees = append(employees, Employee{ID: "102", FirstName: "Ram", LastName: "Bhosale", Salary: 5000.00, Address: &Address{
		StreetName: "Mumbai", PinCode: 354323, State: "Maharashtra", Country: "India",
	}})
	
	// Initialize the router
    router := mux.NewRouter()

    router.HandleFunc("/api/employees", getEmployees).Methods("GET")
	router.HandleFunc("/api/employees/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/api/employees", createEmployee).Methods("POST")
	// router.HandleFunc("/api/employees/{id}", updateEmployee).Methods("PUT")
	// router.HandleFunc("/api/employees/{id}", deleteEmployee).Methods("DELETE")

	http.ListenAndServe(":8090", router)
	log.Println("Server is running on port 8090")

}