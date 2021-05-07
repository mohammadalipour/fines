package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test/controller"
	"test/db"
	"test/repository/mysql"
)

func getEmployee(w http.ResponseWriter, r *http.Request) {
	dbConnection :=db.Connect()

	employeeRepository := mysql.NewEmployeeRepository(dbConnection)
	employeeController :=controller.NewBaseHandler(employeeRepository)
	
	employeeController.Get(w,r)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getEmployee).Methods("GET")

	log.Fatal(http.ListenAndServe(":8002", myRouter))
}

func main() {
	handleRequests()
}
