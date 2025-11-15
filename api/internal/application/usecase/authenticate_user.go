package usecase

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type userAuthenticator struct {
	userService service.UserService
}

func NewUserAuthenticator(
	userService service.UserService,
) contract.UserAuthenticator {
	return &userAuthenticator{
		userService: userService,
	}
}

func (u *userAuthenticator) AuthenticateUser(command *command.AuthenticateUser) (*response.LoginUser, error) {
	user, err := u.userService.FindByParams(*command)
	if err != nil {
		return nil, err
	}

	return &response.LoginUser{
		Token: user.Token,
		User: response.User{
			ID:       user.ID,
			Nickname: user.Nickname,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, nil
}
