package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
)
type Customer struct{
	Name string `json:"name"`
	Gender string `json:"gender"`
	Balance int `json:"balance"`
	AccountNumber int `json:"accountnumber"`
}
var customers []Customer;

func ShowCustomers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(customers)	
	}

	func FindCustomer(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		args := mux.Vars(r);
		var argu = args["id"];
		id,_ := strconv.Atoi(argu)
		for _,data := range customers{
			if(data.AccountNumber == id){
				json.NewEncoder(w).Encode(data)
			}
		}
	}
	func Register(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		var customer Customer;
		json.NewDecoder(r.Body).Decode(&customer)
		customers = append(customers,customer);
		json.NewEncoder(w).Encode(customer);
	}

	func Delete(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		args := mux.Vars(r)
		var argu = args["id"];
		id,_ := strconv.Atoi(argu)
		for index,data := range customers{
			if(data.AccountNumber == id){
             customers = append(customers[:index],customers[index + 1:]...)
			 json.NewEncoder(w).Encode(customers);
			}
		}
	}

	func Update(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		args := mux.Vars(r)
		var argu = args["id"];
		id,_ := strconv.Atoi(argu)
		for index,data := range customers{
			if(data.AccountNumber == id){
             customers = append(customers[:index],customers[index + 1:]...)
			 var customer Customer;
			 json.NewDecoder(r.Body).Decode(&customer)
			 customers = append(customers,customer);
			 json.NewEncoder(w).Encode(customer)
			}
		}
		
	}
func main(){
	/*
	Features --->>>>>>
Read customer details ie name,balance,account number
Show customer details
Add new customer
Find customer by account number
Delete customer details 
Update customer details
	*/
	customers = append(customers,Customer{Name:"Ben San",Gender:"Male",Balance:100,AccountNumber:1234})
	router := mux.NewRouter();
    router.HandleFunc("/ShowAll",ShowCustomers).Methods("GET")
	router.HandleFunc("/Find/{id}",FindCustomer).Methods("GET")
	router.HandleFunc("/Register",Register).Methods("POST")
	router.HandleFunc("/Delete/{id}",Delete).Methods("DELETE")
	router.HandleFunc("/Update/{id}",Update).Methods("PUT")
	fmt.Println("SERVER IS LIVE @ PORT 3000")
	if err := http.ListenAndServe(":3000",router);err != nil{
		log.Fatal(err)
	}
}