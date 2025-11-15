package contract

import (
	"firemap/internal/application/command"
	"firemap/internal/application/response"
)

type UserRegistrator interface {
	RegisterUser(command *command.RegisterUser) (*response.RegisteredUser, error)
}

type UserAuthenticator interface {
	AuthenticateUser(command *command.AuthenticateUser) (*response.LoginUser, error)
}

type UserGetter interface {
	GetUser(token string) (*response.User, error)
}
