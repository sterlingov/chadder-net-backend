package service

import (
	"time"

	httpdto "github.com/sterlingov/chadder-net-backend/internal/delivery/http/dto"
	"github.com/sterlingov/chadder-net-backend/internal/entity"
	"github.com/sterlingov/chadder-net-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *UserService) Register(req *httpdto.CreateUserRequest) (int64, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	user := entity.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}

	if req.Bio != nil {
		user.Bio = *req.Bio
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	return s.repo.Create(&user)
}

func (s *UserService) GetByID(id int64) (*entity.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) GetByUsername(username string) (*entity.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) Update()
