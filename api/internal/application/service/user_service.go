package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"firemap/internal/application/command"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")

type UserService interface {
	Add(user command.RegisterUser) (entity.User, error)
	FindByParams(user command.AuthenticateUser) (entity.User, error)
	FindByToken(token string) (entity.User, error)
}

type userService struct {
	repository contract.UserRepository
}

func NewUserService(repository contract.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) Add(user command.RegisterUser) (entity.User, error) {
	return s.repository.Add(entity.User{
		Email:    user.Email,
		Password: user.Password,
		Nickname: user.Nickname,
		Role:     "user",
		Token:    RandomToken32(),
	})
}

func (s *userService) FindByParams(user command.AuthenticateUser) (entity.User, error) {
	return s.repository.FindByParams(entity.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (s *userService) FindByToken(token string) (entity.User, error) {
	return s.repository.FindByToken(token)
}

func RandomToken32() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
