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
	employeeController :=controller.NewEmployeeBaseHandler(employeeRepository)
	
	employeeController.Get(w,r)
}

func setSalary(w http.ResponseWriter, r *http.Request) {

	dbConnection :=db.Connect()
	levelRepository := mysql.NewLevelRepository(dbConnection)
	levelController :=controller.NewLevelBaseHandler(levelRepository)

	levelController.SetSalary(w,r)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/employee", getEmployee).Methods("GET")
	myRouter.HandleFunc("/level/salary", setSalary).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", myRouter))
}

func main() {
	handleRequests()
}
