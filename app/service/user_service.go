package service

import (
	"errors"
	"time"

	"projectuasbe/app/model"
	"projectuasbe/app/repository"
	"projectuasbe/config"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(username, password string) (*model.UserResponse, error)
	Register(username, email, fullName, password, roleID string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Login(username, password string) (*model.UserResponse, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("username tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("password salah")
	}

	token := generateJWT(user)

	return &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		RoleID:   user.RoleID,
		Token:    token,
	}, nil
}

func (s *userService) Register(username, email, fullName, password, roleID string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		ID:           uuid.NewString(),
		Username:     username,
		Email:        email,
		FullName:     fullName,
		PasswordHash: string(hash),
		RoleID:       roleID,
		IsActive:     true,
	}

	return s.repo.Create(user)
}

func generateJWT(user *model.User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.RoleID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(config.AppConfig.JWTSecret))

	return t
}
