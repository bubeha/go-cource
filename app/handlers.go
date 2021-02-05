package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{{"Test", "City", "ZipCode"}}

	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}