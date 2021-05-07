package mysql

import (
	"database/sql"
	//"log"
	"test/model"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		DB: db,
	}
}

func (r *EmployeeRepository) FindById(id int) (*model.Employee, error) {

	var employee model.Employee
	r.DB.QueryRow("SELECT name,family,level FROM employee where id = ?", id).Scan(&employee.Name, &employee.Family, &employee.Level)

	return &model.Employee{
		Name: employee.Name,
		Family: employee.Family,
		Level: employee.Level,
	}, nil
}

func (r *EmployeeRepository) Save(employee *model.Employee) error {
	return nil
}
