package postgres

import (
	"database/sql"
	"errors"
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

func (r *userRepo) getByField(field string, value any) (*entity.User, error) {
	query := fmt.Sprintf("SELECT id, username, name, bio, created_at, avatar FROM users WHERE %s = $1", field)

	user := &entity.User{}
	var bio sql.NullString
	var avatar sql.NullString

	err := r.db.QueryRow(query, value).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&bio,
		&user.CreatedAt,
		&avatar)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	user.Avatar = avatar.String
	user.Bio = bio.String

	return user, nil
}

func (r *userRepo) GetByID(id int64) (*entity.User, error) {
	return r.getByField("id", id)
}

func (r *userRepo) GetByUsername(username string) (*entity.User, error) {
	return r.getByField("username", username)
}

func (r *userRepo) Update(user *entity.User) (*entity.User, error) {
	query := "UPDATE users SET username = $1, name = $2, email = $3, bio = $4, password_hash = $5, avatar = $6 WHERE id = $7 RETURNING id, username, name, email, bio, password_hash, avatar, created_at"

	updated := &entity.User{}
	var bio, avatar sql.NullString

	err := r.db.QueryRow(query,
		user.Username,
		user.Name,
		user.Email,
		user.Bio,
		user.PasswordHash,
		user.Avatar,
		user.ID,
	).Scan(
		&updated.ID,
		&updated.Username,
		&updated.Name,
		&updated.Email,
		&bio,
		&updated.PasswordHash,
		&avatar,
		&updated.CreatedAt,
	)

	updated.Bio = bio.String
	updated.Avatar = avatar.String

	if err != nil {
		return nil, err
	}
	return updated, nil
}
