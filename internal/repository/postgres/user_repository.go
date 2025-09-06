package postgres

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
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
	query := "INSERT INTO users (name, username, email, bio, password_hash, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var id int64

	err := r.db.QueryRow(query,
		sql.NullString{String: user.Name, Valid: user.Name != ""},
		user.Username,
		user.Email,
		sql.NullString{String: user.Bio, Valid: user.Bio != ""},
		user.PasswordHash,
		user.CreatedAt).Scan(&id)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			switch pgErr.Constraint {
			case "users_email_key":
				return 0, fmt.Errorf("email already exists")
			case "users_username_key":
				return 0, fmt.Errorf("username already exists")
			default:
				return 0, fmt.Errorf("duplicate value")
			}
		}
		return 0, err
	}

	return id, nil
}

func (r *userRepo) GetByID(id int64) (*entity.User, error) {
	query := "SELECT (id, username, name, bio, created_at, avatar) FROM users WHERE id = $1"

	user := &entity.User{}
	var bio sql.NullString
	var avatar sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&bio,
		&user.CreatedAt,
		&avatar)

	if err != nil {
		return nil, err
	}
	user.Avatar = avatar.String
	user.Bio = bio.String

	return user, nil
}

func (r *userRepo) GetByUsername(username string) (*entity.User, error) {
	//Затычка
	return &entity.User{}, nil
}
