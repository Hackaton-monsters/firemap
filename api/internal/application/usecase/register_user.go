package usecase

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type userRegistrator struct {
	userService service.UserService
}

func NewUserRegistrator(
	userService service.UserService,
) contract.UserRegistrator {
	return &userRegistrator{
		userService: userService,
	}
}

func (u *userRegistrator) RegisterUser(command *command.RegisterUser) (*response.RegisteredUser, error) {
	user, err := u.userService.Add(*command)
	if err != nil {
		return nil, err
	}

	return &response.RegisteredUser{
		Token: user.Token,
		User: response.User{
			ID:       user.ID,
			Nickname: user.Nickname,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, nil
}
