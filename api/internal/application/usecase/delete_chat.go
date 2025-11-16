package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/service"
)

type chatDeleter struct {
	userService     service.UserService
	chatUserService service.ChatUserService
}

func NewChatDeleter(
	userService service.UserService,
	chatUserService service.ChatUserService,
) contract.ChatDeleter {
	return &chatDeleter{
		userService:     userService,
		chatUserService: chatUserService,
	}
}

func (u *chatDeleter) DeleteChat(token string, chatID int64) error {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return err
	}

	err = u.chatUserService.Delete(user.ID, chatID)
	if err != nil {
		return err
	}

	return nil
}
