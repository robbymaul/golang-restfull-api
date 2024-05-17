package repository

import (
	"github.com/robbymaul/golang-swagger.git/helper"
	"github.com/robbymaul/golang-swagger.git/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Register(user model.User) (model.User, error) {
	tx := r.DB.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}

	if err := tx.Commit().Error; err != nil {
		helper.PanicIfError(err)
	}

	return user, nil
}

func (r *UserRepositoryImpl) Login(username string) (model.User, error) {
	var user model.User

	err := r.DB.Where("username", username).First(&user).Error
	if err != nil {
		helper.PanicIfError(err)
	}

	return user, nil
}
