package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/model"
)

type BaseHandler struct {
	employeeRepository model.EmployeeRepository
}

func NewBaseHandler(repository model.EmployeeRepository)*BaseHandler  {
	return &BaseHandler{
		employeeRepository: repository,
	}
}

func (baseHandler *BaseHandler) Get(w http.ResponseWriter, r *http.Request)  {
	 Id, err := strconv.Atoi(r.FormValue("id"))

	employee, err := baseHandler.employeeRepository.FindById(Id)

	if err != nil {
		fmt.Println("Error", employee)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}