package service

import (
	"errors"
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

func userToUserResponse(user *entity.User) *httpdto.UserResponse {
	userResp := httpdto.UserResponse{}
	userResp.ID = user.ID
	userResp.Name = user.Name
	userResp.Username = user.Username
	userResp.Bio = user.Bio
	userResp.Avatar = user.Avatar
	userResp.CreatedAt = user.CreatedAt.Format(time.RFC3339)
	return &userResp
}

func (s *UserService) GetByID(id int64) (*httpdto.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *UserService) GetByUsername(username string) (*httpdto.UserResponse, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *UserService) Update(userID int64, req *httpdto.UpdateUserRequest) (*httpdto.UserResponse, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Bio != nil {
		user.Bio = *req.Bio
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Password != nil {
		var hash string
		hash, err = hashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hash
	}

	var newUser *entity.User
	newUser, err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return userToUserResponse(newUser), nil
}
