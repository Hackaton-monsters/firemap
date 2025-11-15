package repository

import (
	"errors"
	"firemap/internal/application/service"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) contract.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Add(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return user, service.ErrUserAlreadyExists
		}
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByParams(params entity.User) (entity.User, error) {
	var user entity.User
	err := r.db.
		Where("users.email = ?", params.Email).
		Where("users.password = ?", params.Password).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, service.ErrUserNotFound
		}
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByToken(token string) (entity.User, error) {
	var user entity.User
	err := r.db.
		Where("users.token = ?", token).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, service.ErrUserNotFound
		}
		return user, err
	}

	return user, nil
}
