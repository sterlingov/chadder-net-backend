package postgres

import (
	"database/sql"

	"github.com/sterlingov/chadder-net-backend/internal/entity"
	"github.com/sterlingov/chadder-net-backend/internal/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *entity.User) (int64, error) {
	//Затычка
	return 0, nil
}

func (r *userRepo) GetByID(id int64) (*entity.User, error) {
	//Затычка
	return &entity.User{}, nil
}

func (r *userRepo) GetByUsername(username string) (*entity.User, error) {
	//Затычка
	return &entity.User{}, nil
}
