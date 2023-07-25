package repository

import (
	"golang_clean_architecture_sample/model"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
