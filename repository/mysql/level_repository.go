package mysql

import (
	"database/sql"
	"fmt"
	"test/model"
)

type LevelRepository struct {
	DB *sql.DB
}

func NewLevelRepository(db *sql.DB) *LevelRepository {
	return &LevelRepository{
		DB: db,
	}
}

func (r *LevelRepository) FindById(id int) (*model.Level, error) {

	var level model.Level
	r.DB.QueryRow("SELECT title,salary,is_active"+
		" FROM levels e "+
		" where id = ?", id).Scan(&level.Title, &level.Salary, &level.IsActive)
	defer r.DB.Close()

	return &model.Level{
		Title:  level.Title,
		Salary: level.Salary,
		IsActive: level.IsActive,
	}, nil
}

func (r *LevelRepository) Save(level *model.Level) error {
	_, err := r.DB.Exec("update levels set salary=? where title = ?", level.Salary, level.Title)
	defer r.DB.Close()

	if err != nil {
		fmt.Println(err)
	}

	return err
}
