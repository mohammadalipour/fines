package model

type Level struct {
	Title    string
	Salary   int
	IsActive bool
}

type LevelRepository interface {
	FindById(id int) (*Level, error)
	Save(level *Level) error
}