package repository

import "github.com/sterlingov/chadder-net-backend/internal/entity"

type UserRepository interface {
	Create(user *entity.User) (int64, error)
	GetByID(id int64) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}
