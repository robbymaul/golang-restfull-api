package repository

import (
	"github.com/robbymaul/golang-swagger.git/model"
)

type UserRepository interface {
	Register(user model.User) (model.User, error)
	Login(username string) (model.User, error)
}
