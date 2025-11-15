package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type userGetter struct {
	userService service.UserService
}

func NewUserGetter(
	userService service.UserService,
) contract.UserGetter {
	return &userGetter{
		userService: userService,
	}
}

func (u *userGetter) GetUser(token string) (*response.User, error) {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:       user.ID,
		Nickname: user.Nickname,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}
