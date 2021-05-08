package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/model"
)

type LevelBaseHandler struct {
	levelRepository model.LevelRepository
}

func NewLevelBaseHandler(repository model.LevelRepository) *LevelBaseHandler {
	return &LevelBaseHandler{
		levelRepository: repository,
	}
}

func (baseHandler *LevelBaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.Atoi(r.FormValue("id"))

	employee, err := baseHandler.levelRepository.FindById(Id)

	if err != nil {
		fmt.Println("Error", employee)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

func (baseHandler *LevelBaseHandler) SetSalary(w http.ResponseWriter, r *http.Request) {
	Title := r.FormValue("level")
	Salary, err := strconv.Atoi(r.FormValue("salary"))

	levelStruct := &model.Level{
		Title:    Title,
		Salary:   Salary,
		IsActive: true,
	}

	err = baseHandler.levelRepository.Save(levelStruct)

	if err != nil {
		fmt.Println("Error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("salary updated")
}
