package model

type Employee struct {
	Name   string
	Family string
	Level  string
	Salary int
}

type EmployeeRepository interface {
	FindById(id int) (*Employee, error)
	Save(user *Employee) error
}