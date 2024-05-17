package service

import (
	"github.com/robbymaul/golang-swagger.git/model"
	repository "github.com/robbymaul/golang-swagger.git/repository/user"
	web "github.com/robbymaul/golang-swagger.git/web"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) Register(request *web.UserRegisterRequest) (web.UserRegisterResponse, error) {
	user := model.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := service.UserRepository.Register(user)
	if err != nil {
		return web.UserRegisterResponse{}, err
	}

	return web.UserRegisterResponse{
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (service *UserServiceImpl) Login(request web.UserLoginRequest) (web.UserLoginResponse, error) {

	return web.UserLoginResponse{}, nil
}
