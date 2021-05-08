package mysql

import (
	"database/sql"
	"fmt"
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
	r.DB.QueryRow("SELECT e.name,e.family,l.title level,(ehw.hours_worked * l.salary) salary"+
		" FROM employees e "+
		" INNER JOIN levels l ON l.id=e.level_id "+
		" INNER JOIN employee_hours_worked ehw on e.id = ehw.employee_id"+
		" where e.id = ?", id).Scan(&employee.Name, &employee.Family, &employee.Level,&employee.Salary)
	defer r.DB.Close()

	fmt.Println(employee)
	return &model.Employee{
		Name:   employee.Name,
		Family: employee.Family,
		Level:  employee.Level,
		Salary: employee.Salary,
	}, nil
}

func (r *EmployeeRepository) Save(employee *model.Employee) error {
	return nil
}
