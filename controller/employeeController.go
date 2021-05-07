package controller

import (
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

	fmt.Println("result: ",employee)
}